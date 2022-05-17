package auth

import (
	"github.com/alekstet/question/conf"

	_ "github.com/mattn/go-sqlite3"
)

type S conf.Store

func New(store conf.Store) *S {
	return &S{
		Db:      store.Db,
		Log:     store.Log,
		Routes:  store.Routes,
		Session: store.Session,
	}
}

func (s *S) Register() {
	s.Routes.POST("/signup", s.SignUp)
	s.Routes.POST("/signin", s.SignIn)
	s.Routes.POST("/logout", s.Logout)
}
