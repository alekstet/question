package routes

import (
	"net/http"
	"question/api/routes/auth"
	"question/api/routes/process"
	"question/conf"

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

func Routes(s conf.Store) *httprouter.Router {
	process.New(s).Register()
	auth.New(s).Register()
	return s.Routes
}

func (rr *R) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rr.Router.ServeHTTP(w, r)
}
