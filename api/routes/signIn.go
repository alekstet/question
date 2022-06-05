package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/alekstet/question/api/errors"
	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

const existUser = `
SELECT COUNT(*), Password, Nickname FROM users_auth
WHERE Login = $1`

func (s *Store) signIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		data           models.SignIn
		nickname, hash string
		exists         int
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	if !data.Valid() {
		helpers.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = s.Db.QueryRow(existUser, data.Login).Scan(&exists, &hash, &nickname)
	if err != nil {
		helpers.Error(w, r, http.StatusUnauthorized, errors.ErrIncorectAuthData)
		return
	}

	if exists != 1 && !helpers.CheckPasswordHash(data.Password, hash) {
		helpers.Error(w, r, http.StatusUnauthorized, errors.ErrIncorectAuthData)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Login: data.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	err = godotenv.Load()
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	jwtKey := os.Getenv("JWTKey")
	jwtKeyByte := []byte(jwtKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeyByte)
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
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}

		helpers.Error(w, r, http.StatusBadRequest, err)
		return
	}

	tknStr := cookie.Value
	claims := &models.Claims{}

	err = godotenv.Load()
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	jwtKey := os.Getenv("JWTKey")
	jwtKeyByte := []byte(jwtKey)

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKeyByte, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		helpers.Error(w, r, http.StatusBadRequest, err)
		return
	}

	if !tkn.Valid {
		helpers.Error(w, r, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Login)))
}
