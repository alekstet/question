package routes

import (
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) Register() {
	s.Routes.POST("/answer", s.createAnswer)
	s.Routes.PUT("/answer", s.updateAnswer)
	s.Routes.DELETE("/answer", s.deleteAnswer)
	s.Routes.GET("/new", s.getTodayAnswers)
	s.Routes.GET("/users", s.getUsers)
	s.Routes.GET("/users/:user", s.getUserInfo)
	s.Routes.POST("/question", s.createQuestion)
	s.Routes.POST("/signin", s.signIn)
	s.Routes.POST("/signup", s.signUp)
	s.Routes.POST("/welcome", s.welcome)
}
