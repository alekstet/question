package routes

import (
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) Register() {
	s.Routes.PUT("/new", s.updateAnswer)
	s.Routes.DELETE("/new", s.deleteAnswer)
	s.Routes.GET("/new", s.getTodayAnswers)
	s.Routes.GET("/users", s.getUsers)
	s.Routes.GET("/users/:user", s.getUserInfo)
	s.Routes.POST("/new", s.createAnswer)
	s.Routes.POST("/questions", s.addQuestion)
	s.Routes.POST("/signin", s.signIn)
	s.Routes.POST("/signup", s.signUp)
	s.Routes.POST("/welcome", s.welcome)
}
