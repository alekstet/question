package routes

import (
	"net/http"

	"github.com/alekstet/question/helpers"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) getTodayAnswers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page := r.URL.Query().Get("page")

	todaysInfo, err := s.Querier.GetTodayAnswers(page)
	if err != nil {
		helpers.Render(w, r, http.StatusInternalServerError, nil)
		return
	}

	helpers.Render(w, r, http.StatusOK, todaysInfo)
}
