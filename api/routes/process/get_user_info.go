package process

import (
	"fmt"
	"net/http"

	model "github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) UserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var name string
	var sex string
	userAnsw := []model.UserAnsw{}
	sort := r.URL.Query().Get("sort")
	nickname := params.ByName("user")

	sql := fmt.Sprintf(
		`SELECT Question_id, Question, Answer FROM users_questions 
		INNER JOIN questions ON Question_id = Date WHERE User_nickname = '%s'`, nickname)
	if sort == "dateup" {
		sql = fmt.Sprintf("%s ORDER BY Date ASC", sql)
	}
	if sort == "datedown" {
		sql = fmt.Sprintf("%s ORDER BY Date DESC", sql)
	}
	rows, err := s.Db.Query(sql)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := model.UserAnsw{}
		err := rows.Scan(&p.Date, &p.Question, &p.Answer)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}
		userAnsw = append(userAnsw, p)
	}

	err = s.Db.QueryRow("SELECT Name, Sex FROM users_data WHERE User_nickname = $1", nickname).Scan(&name, &sex)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}

	userInfo := model.UserInfo{Name: name, Sex: sex, Answers: userAnsw}
	helpers.Render(w, r, 200, userInfo)
}
