package process

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"question/api/models"
	"question/helpers"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) CreateAnswer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var exists int
	date := time.Now().Format("02.01.2006 15:04:05")
	data := &models.User_question{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)

	err = s.Db.QueryRow(
		`SELECT EXISTS
		(SELECT * FROM users_questions 
			WHERE Question_id = $1 AND User_nickname = $2)`, data.QuestionId, data.UserNickname).Scan(&exists)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}

	if exists == 1 {
		helpers.Error(w, r, 400, errors.New("user already exists"))
		return
	} else {
		_, err = s.Db.Exec(
			`INSERT INTO users_questions (Question_id, User_nickname, Answer, Created_at, Updated_at) 
			VALUES ($1, $2, $3, $4, $5)`, data.QuestionId, data.UserNickname, data.Answer, date, date)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}
	}

	helpers.Render(w, r, 201, nil)
}
