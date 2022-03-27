package auth

import (
	"github.com/alekstet/question/conf"

	_ "github.com/mattn/go-sqlite3"
)

type S conf.Store

func New(s conf.Store) *S {
	return &S{
		Db:      s.Db,
		Log:     s.Log,
		Routes:  s.Routes,
		Session: s.Session,
	}
}

func (s *S) Register() {
	s.Routes.POST("/signup", s.SignUp)
	s.Routes.POST("/signin", s.SignIn)
	s.Routes.POST("/logout", s.Logout)
}
