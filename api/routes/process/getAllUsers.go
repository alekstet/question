package process

import (
	"fmt"
	"net/http"

	"github.com/alekstet/question/helpers"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Im here")
	fmt.Println(s.Querier)
	users, err := s.Querier.GetUsers()
	if err != nil {
		helpers.Render(w, r, http.StatusInternalServerError, nil)
		return
	}

	helpers.Render(w, r, http.StatusOK, users)
}
