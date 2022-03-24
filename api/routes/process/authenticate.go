package process

import (
	"fmt"
	"net/http"
	"question/helpers"

	"github.com/julienschmidt/httprouter"
)

const session_name = "current_session"

func (s *S) Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}

func (s *S) Authenticate(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session, err := s.Session.Get(r, session_name)
		if err != nil {
			helpers.Error(w, r, 500, err)
			return
		}

		fmt.Println("VAL:", session.Values)

		nickname, ok := session.Values["user_nickname"]

		if !ok {
			helpers.Error(w, r, 400, err)
		} else {
			fmt.Println(nickname)
			h(w, r, ps)
		}
	}
}
