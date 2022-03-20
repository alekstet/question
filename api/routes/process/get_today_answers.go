package process

import (
	"fmt"
	"net/http"
	model "question/api/models"
	"question/helpers"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) TodayAnswers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	date := time.Now().Format("02.01.2006")
	answers := []model.TodaysAnswer{}
	var todaysQuestion string
	var page int
	perPage := 2

	sql := fmt.Sprintf(
		`SELECT Nickname, Answer FROM users_auth 
		INNER JOIN users_questions ON Nickname = User_nickname 
		INNER JOIN questions ON Date = Question_id WHERE Date = '%s'`, date)
	if r.URL.Query().Get("page") == "" {
		sql = fmt.Sprintf("%s LIMIT %v", sql, perPage)
	}
	if r.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		sql = fmt.Sprintf("%s LIMIT %v OFFSET %v", sql, perPage, (page-1)*perPage)
	}

	rows, err := s.Db.Query(sql)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := model.TodaysAnswer{}
		err := rows.Scan(&p.Nickname, &p.Answer)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}
		answers = append(answers, p)
	}

	err = s.Db.QueryRow("SELECT Question FROM questions WHERE Date = $1", date).Scan(&todaysQuestion)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}

	todaysInfo := model.TodaysInfo{Question: todaysQuestion, Answers: answers}
	helpers.Render(w, r, 200, todaysInfo)
}
