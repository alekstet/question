package routes

import (
	"net/http"

	"github.com/alekstet/question/api/routes/process"
	"github.com/julienschmidt/httprouter"
)

type R struct {
	Router *httprouter.Router
}

func New() *R {
	return &R{
		Router: httprouter.New(),
	}
}

func Routes(s process.Store) *httprouter.Router {
	s.Register()
	return s.Routes
}

func (rr *R) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rr.Router.ServeHTTP(w, r)
}
