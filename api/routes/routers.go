package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type R struct {
	Router *httprouter.Router
}

func NewRouter() *R {
	return &R{
		Router: httprouter.New(),
	}
}

func Routes(s Store) *httprouter.Router {
	s.Register()
	return s.Routes
}

func (rr *R) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rr.Router.ServeHTTP(w, r)
}
