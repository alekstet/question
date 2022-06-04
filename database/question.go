package database

import (
	"time"

	"github.com/alekstet/question/api/errors"
	"github.com/alekstet/question/api/models"
)

const (
	questionExists = `
SELECT EXISTS
(SELECT * FROM questions 
WHERE Date = $1 AND Question = $2)`
	insertQuestion = `
INSERT INTO questions (Date, Question) 
VALUES ($1, $2)`
)

func (s *Store) AddQuestion(data models.Question) error {
	var exists int
	date := time.Now().Format("02.01.2006")
	if !data.Valid() {
		return errors.ErrDataNotValid
	}

	err := s.Db.QueryRow(questionExists, date, data.Question).Scan(&exists)
	if err != nil {
		return err
	}

	if exists == 1 {
		return err
	}

	_, err = s.Db.Exec(insertQuestion, date, data.Question)
	if err != nil {
		return err
	}

	return nil
}
