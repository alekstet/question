package database

import (
	"github.com/alekstet/question/api/errors"
	"github.com/alekstet/question/api/models"
	"github.com/alekstet/question/helpers"
)

const (
	existUser = `
SELECT COUNT(*), Password, Nickname FROM users_auth
WHERE Login = $1`

	insertUsersCreds = `
INSERT INTO users_auth (Login, Password, Nickname) VALUES ($1, $2, $3)`
)

func (s Store) SignUp(data models.SignUp) error {
	err := data.Valid()
	if err != nil {
		return err
	}

	encryptedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		return err
	}

	_, err = s.Db.Exec(insertUsersCreds, data.Login, encryptedPassword, data.Nickname)
	if err != nil {
		return err
	}

	return nil
}

func (s Store) SignIn(data models.SignIn) (string, error) {
	var (
		nickname, hash string
		exists         int
	)

	if !data.Valid() {
		err := errors.ErrDataNotValid
		return "", err
	}

	err := s.Db.QueryRow(existUser, data.Login).Scan(&exists, &hash, &nickname)
	if err != nil {
		return "", err
	}

	if exists != 1 && !helpers.CheckPasswordHash(data.Password, hash) {
		err = errors.ErrIncorectAuthData
		return "", err
	}

	return data.Login, nil
}
