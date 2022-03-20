package process

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"question/api/models"
	"question/helpers"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) UpdateAnswer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	updated_at := time.Now().Format("02.01.2006 15:04:05")
	data := &models.User_question{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)

	row, err := s.Db.Exec(
		`UPDATE users_questions SET Answer = $1, Updated_at = $2
		WHERE Question_id = $3 AND User_nickname = $4`,
		data.Answer, updated_at, data.QuestionId, data.UserNickname)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	isUpdated, _ := row.RowsAffected()
	if isUpdated == 0 {
		helpers.Error(w, r, 500, err)
		return
	}

	helpers.Render(w, r, 200, nil)
}
