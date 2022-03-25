package process

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	er "question/api/errors"
	"question/api/models"
	"question/helpers"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) AddQuestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var exists int
	date := time.Now().Format("02.01.2006")
	data := &models.Question{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)

	if !data.Valid() {
		helpers.Error(w, r, 400, er.ErrDataNotValid)
		return
	}

	err = s.Db.QueryRow(
		`SELECT EXISTS
		(SELECT * FROM questions 
			WHERE Date = $1 AND Question = $2)`, date, data.Question).Scan(&exists)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}

	if exists == 1 {
		helpers.Error(w, r, 400, errors.New("question already exists"))
		return
	} else {
		_, err = s.Db.Exec(
			`INSERT INTO questions (Date, Question) 
			VALUES ($1, $2)`, date, data.Question)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}
	}

	helpers.Render(w, r, 201, nil)
}
