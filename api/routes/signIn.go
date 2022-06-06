package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"
	"github.com/julienschmidt/httprouter"
)

func (s *Store) signIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data models.SignIn

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

	login, err := s.Querier.SignIn(data)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	/* err = godotenv.Load()
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	jwtKey := os.Getenv("JWTKey") */

	signInData, err := s.TokenMaker.CreateToken(login)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signInData.Token,
		Expires: signInData.ExpTime,
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

	claims, err := s.TokenMaker.VerifyToken(tknStr)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Login)))
}
