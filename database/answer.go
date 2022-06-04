package database

import (
	"time"

	"github.com/alekstet/question/api/errors"
	"github.com/alekstet/question/api/models"
)

func (s Store) CreateAnswer(data models.UserQuestion) error {
	var exists int
	timeNow := time.Now().Format("02.01.2006 15:04:05")

	if !data.Valid() {
		return errors.ErrDataNotValid
	}

	err := s.Db.QueryRow(
		`SELECT EXISTS
		(SELECT * FROM users_questions 
			WHERE Question_id = $1 AND User_nickname = $2)`, data.QuestionId, data.UserNickname).Scan(&exists)
	if err != nil {
		return err
	}

	if exists == 1 {
		return err
	} else {
		_, err = s.Db.Exec(
			`INSERT INTO users_questions (Question_id, User_nickname, Answer, Created_at, Updated_at) 
			VALUES ($1, $2, $3, $4, $5)`, data.QuestionId, data.UserNickname, data.Answer, timeNow, timeNow)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Store) UpdateAnswer(data models.UserQuestion) error {
	updatedAt := time.Now().Format("02.01.2006 15:04:05")
	row, err := s.Db.Exec(
		`UPDATE users_questions SET Answer = $1, Updated_at = $2
		WHERE Question_id = $3 AND User_nickname = $4`,
		data.Answer, updatedAt, data.QuestionId, data.UserNickname)
	if err != nil {
		return err
	}

	isUpdated, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if isUpdated == 0 {
		return err
	}

	return nil
}

func (s Store) DeleteAnswer(data models.UserQuestion) error {
	_, err := s.Db.Exec(
		`DELETE FROM users_questions 
		WHERE Question_Id = $1 AND User_nickname = $2`, data.QuestionId, data.UserNickname)
	if err != nil {
		return err
	}

	return nil
}
