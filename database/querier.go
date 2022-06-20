package database

import (
	"context"

	"github.com/alekstet/question/api/models"
)

type Querier interface {
	UpdateAnswer(context.Context, models.UserQuestion) error
	CreateAnswer(context.Context, models.UserQuestion) error
	DeleteAnswer(context.Context, models.UserQuestion) error
	GetUsers(context.Context) ([]models.UsersData, error)
	AddQuestion(context.Context, models.Question) error
	GetUserInfo(context.Context, string, string) (*models.UserInfo, error)
	GetTodayAnswers(context.Context, string) (*models.TodaysInfo, error)
	SignIn(context.Context, models.SignIn) (string, error)
	SignUp(context.Context, models.SignUp) error
}

var _ Querier = (*Store)(nil)
