package process

import (
	"net/http"

	model "github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := s.Db.Query("SELECT User_nickname, Name, Sex FROM users_data")
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	defer rows.Close()

	users := []model.UsersData{}

	for rows.Next() {
		p := model.UsersData{}
		err := rows.Scan(&p.UserNickname, &p.Name, &p.Sex)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}
		users = append(users, p)
	}

	helpers.Render(w, r, 200, users)
}
