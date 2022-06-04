package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alekstet/question/api/errors"
	"github.com/alekstet/question/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *Store) signIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		helpers.Error(w, r, http.StatusBadRequest, err)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		err := errors.ErrIncorectAuthData
		helpers.Error(w, r, http.StatusUnauthorized, err)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func (s *Store) welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		helpers.Error(w, r, http.StatusBadRequest, err)
		return
	}

	tknStr := c.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
