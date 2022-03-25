package auth

import (
	"net/http"
	"question/helpers"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func (s *S) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := s.Session.Get(r, session_name)
	if err != nil {
		helpers.Error(w, r, 500, err)
		return
	}
	delete(session.Values, "user_nickname")
	session.Save(r, w)
}
