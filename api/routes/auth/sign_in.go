package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"question/api/models"
	"question/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

const session_name = "current_session"

func (s *S) SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data := &models.SignIn{}
	var nick string
	var hash string
	var exists int

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	json.Unmarshal(body, &data)

	err = s.Db.QueryRow(
		`SELECT COUNT(*), Password, Nickname FROM users_auth
			WHERE Login = $1`, data.Login).Scan(&exists, &hash, &nick)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}

	session, _ := s.Session.Get(r, session_name)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}

	session.Values["user_nickname"] = nick
	err = s.Session.Save(r, w, session)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}

	if exists == 1 && helpers.CheckPasswordHash(data.Password, hash) {
		helpers.Render(w, r, 200, nil)
		return
	} else {
		helpers.Error(w, r, 401, errors.New("login or password are incorrect"))
		return
	}
}
