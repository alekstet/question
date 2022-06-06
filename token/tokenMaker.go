package token

import "github.com/alekstet/question/api/models"

type TokenMaker interface {
	CreateToken(login string) (*models.SignInData, error)
	VerifyToken(token string) (*models.Claims, error)
}

var _ TokenMaker = (*JWTTokenMaker)(nil)
