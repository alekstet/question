package models

import (
	"errors"
	"fmt"
)

type SignUp struct {
	Nickname        string `json:"nickname"`
	Login           string `json:"login"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confPassword"`
}

type SignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (su *SignUp) Init(nickname, login, password, confPassword string) {
	su.Nickname = nickname
	su.Login = login
	su.Password = password
	su.ConfirmPassword = confPassword
}

func (su *SignIn) Init(login, password string) {
	su.Login = login
	su.Password = password
}

func (su *SignUp) Valid() error {
	if len(su.Login) < 3 {
		return fmt.Errorf("len of Login must greater than 2 characters, your len: %v", len(su.Login))
	}
	if len(su.Password) < 2 {
		return fmt.Errorf("len of Password must greater than 2 characters, your len: %v", len(su.Password))
	}
	if su.Password != su.ConfirmPassword {
		return errors.New("passwords dont match")
	}
	return nil
}
