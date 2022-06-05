package database

import (
	"github.com/alekstet/question/api/models"
)

type Querier interface {
	UpdateAnswer(models.UserQuestion) error
	CreateAnswer(models.UserQuestion) error
	DeleteAnswer(models.UserQuestion) error
	GetUsers() ([]models.UsersData, error)
	AddQuestion(data models.Question) error
	GetUserInfo(nickname, sort string) (*models.UserInfo, error)
	GetTodayAnswers(page string) (*models.TodaysInfo, error)
}

var _ Querier = (*Store)(nil)
