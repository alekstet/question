package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data := &models.SignUp{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)

	err = data.Valid()
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}

	fmt.Println(data)

	encrypted_password, err := helpers.HashPassword(data.Password)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}

	_, err = s.Db.Exec(
		`INSERT INTO users_auth (Login, Password, Nickname) VALUES ($1, $2, $3)`,
		data.Login, encrypted_password, data.Nickname)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}

	helpers.Render(w, r, 201, nil)
}
