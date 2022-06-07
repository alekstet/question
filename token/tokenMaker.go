package token

import (
	"time"

	"github.com/alekstet/question/api/models"
)

type TokenMaker interface {
	CreateToken(string, time.Duration) (*models.SignInData, error)
	VerifyToken(string) (*models.Claims, error)
}

var _ TokenMaker = (*JWTTokenMaker)(nil)
