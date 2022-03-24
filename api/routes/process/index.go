package process

import (
	"question/conf"

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
	//s.Routes.POST("/new", s.CreateAnswer)
	s.Routes.PUT("/new", s.UpdateAnswer)
	s.Routes.DELETE("/new", s.DeleteAnswer)
	s.Routes.GET("/new", s.TodayAnswers)
	s.Routes.GET("/users", s.GetUsers)
	s.Routes.GET("/users/:user", s.UserInfo)
	s.Routes.POST("/new", s.Authenticate(s.CreateAnswer))
	s.Routes.POST("/questions", s.AddQuestion)
}
