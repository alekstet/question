package errors

import "errors"

var (
	ErrDataNotValid     = errors.New("data not valid")
	ErrExistsAnswer     = errors.New("answer of this user already exists")
	ErrQuestionExists   = errors.New("question already exists")
	ErrIncorectAuthData = errors.New("login or password are incorrect")
)
