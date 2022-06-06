package token

import (
	"os"
	"time"

	"github.com/alekstet/question/api/errors"
	"github.com/alekstet/question/api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type JWTTokenMaker struct {
	SymmetricKey string
}

func (maker *JWTTokenMaker) CreateToken(login, jwtKey string) (*models.SignInData, error) {
	expTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}

	signInData := &models.SignInData{
		Token:   token,
		ExpTime: expTime,
	}

	return signInData, nil
}

func (maker *JWTTokenMaker) VerifyToken(token string) (*models.Claims, error) {
	claims := &models.Claims{}
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		jwtKey := os.Getenv("JWTKey")
		jwtKeyByte := []byte(jwtKey)
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
