package process

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) DeleteAnswer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	data := &models.User_question{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, r, 400, err)
		return
	}
	json.Unmarshal(body, &data)
	_, err = s.Db.Exec(
		`DELETE FROM users_questions 
		WHERE Question_Id = $1 AND User_nickname = $2`, data.QuestionId, data.UserNickname)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}

	helpers.Render(w, r, 200, nil)
}
