package token

import (
	"time"

	"github.com/alekstet/question/api/errors"
	"github.com/alekstet/question/api/models"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenMaker struct {
	SymmetricKey string
}

func NewPayload(login string, duration time.Duration) *models.Claims {
	claims := &models.Claims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	return claims
}

func (maker *JWTTokenMaker) CreateToken(login string, duration time.Duration) (*models.SignInData, error) {
	claims := NewPayload(login, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(maker.SymmetricKey))
	if err != nil {
		return nil, err
	}

	signInData := &models.SignInData{
		Token:   token,
		ExpTime: time.Now().Add(duration),
	}

	return signInData, nil
}

func (maker *JWTTokenMaker) VerifyToken(token string) (*models.Claims, error) {
	claims := &models.Claims{}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		jwtKeyByte := []byte(maker.SymmetricKey)
		return jwtKeyByte, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, claims, keyFunc)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}

		return nil, err
	}

	if !jwtToken.Valid {
		err := errors.ErrTokenNotValid
		return nil, err
	}

	return claims, nil
}
