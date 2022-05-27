package process

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) getTodayAnswers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var todaysQuestion string
	perPage := 2
	timeNow := time.Now().Format("02.01.2006")
	answers := []models.TodaysAnswer{}

	sql := fmt.Sprintf(
		`SELECT Nickname, Answer FROM users_auth 
		INNER JOIN users_questions ON Nickname = User_nickname 
		INNER JOIN questions ON Date = Question_id WHERE Date = '%s'`, timeNow)

	if r.URL.Query().Get("page") == "" {
		sql = fmt.Sprintf("%s LIMIT %v", sql, perPage)
	}

	if r.URL.Query().Get("page") != "" {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			helpers.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		sql = fmt.Sprintf("%s LIMIT %v OFFSET %v", sql, perPage, (page-1)*perPage)
	}

	rows, err := s.Db.Query(sql)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		todaysAnswer := models.TodaysAnswer{}
		err := rows.Scan(&todaysAnswer.Nickname, &todaysAnswer.Answer)
		if err != nil {
			helpers.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		answers = append(answers, todaysAnswer)
	}

	err = s.Db.QueryRow("SELECT Question FROM questions WHERE Date = $1", timeNow).Scan(&todaysQuestion)
	if err != nil {
		helpers.Error(w, r, http.StatusInternalServerError, err)
		return
	}

	todaysInfo := models.TodaysInfo{Question: todaysQuestion, Answers: answers}
	helpers.Render(w, r, http.StatusOK, todaysInfo)
}
