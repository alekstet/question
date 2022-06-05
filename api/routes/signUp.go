package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"
	"github.com/julienschmidt/httprouter"
)

func (s *Store) signUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data models.SignUp
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}
	json.Unmarshal(body, &data)

	err = data.Valid()
	if err != nil {
		helpers.Error(w, r, http.StatusBadRequest, err)
		return
	}

	encryptedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	_, err = s.Db.Exec(
		`INSERT INTO users_auth (Login, Password, Nickname) VALUES ($1, $2, $3)`,
		data.Login, encryptedPassword, data.Nickname)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	helpers.Render(w, r, http.StatusCreated, nil)
}
