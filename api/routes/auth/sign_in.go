package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"question/api/models"
	"question/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var exists int
	var hash interface{}
	data := &models.SignIn{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	json.Unmarshal(body, &data)

	/* err = s.Db.QueryRow(
	`SELECT EXISTS
	(SELECT * FROM users_auth
		WHERE Login = $1 AND Password = $2)`, data.Login, data.Password).Scan(&exists) */

	err = s.Db.QueryRow(
		`SELECT COUNT(*), Password FROM users_auth
			WHERE Login = $1 AND Password = $2`, data.Login, data.Password).Scan(&exists, &hash)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	fmt.Println(hash)

	if exists == 1 && helpers.CheckPasswordHash(data.Password, hash.(string)) {
		helpers.Render(w, r, 200, nil)
		return
	} else {
		helpers.Error(w, r, 401, errors.New("login or password are incorrect"))
		return
	}
}
