package routes

import (
	"net/http"

	"github.com/alekstet/question/helpers"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) getUserInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	sort := r.URL.Query().Get("sort")
	nickname := params.ByName("user")

	userInfo, err := s.Querier.GetUserInfo(nickname, sort)
	if err != nil {
		helpers.Render(w, r, http.StatusInternalServerError, userInfo)
		return
	}

	helpers.Render(w, r, http.StatusOK, userInfo)
}
