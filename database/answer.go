package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/alekstet/question/api/errors"
	"github.com/alekstet/question/api/models"
)

const getAnswers = `
SELECT Nickname, Answer FROM users_auth 
INNER JOIN users_questions ON Nickname = User_nickname 
INNER JOIN questions ON Date = Question_id WHERE Date = '%s'`

func (s Store) GetTodayAnswers(ctx context.Context, page string) (*models.TodaysInfo, error) {
	perPage := 2
	timeNow := time.Now().Format("02.01.2006")
	answers := []models.TodaysAnswer{}

	sql := fmt.Sprintf(getAnswers, timeNow)

	if page == "" {
		sql = fmt.Sprintf("%s LIMIT %v", sql, perPage)
	}

	if page != "" {
		page, err := strconv.Atoi(page)
		if err != nil {
			return nil, err
		}

		sql = fmt.Sprintf("%s LIMIT %v OFFSET %v", sql, perPage, (page-1)*perPage)
	}

	rows, err := s.Db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		todaysAnswer := models.TodaysAnswer{}
		err := rows.Scan(&todaysAnswer.Nickname, &todaysAnswer.Answer)
		if err != nil {
			return nil, err
		}

		answers = append(answers, todaysAnswer)
	}

	var todaysQuestion string

	err = s.Db.QueryRowContext(ctx, "SELECT Question FROM questions WHERE Date = $1", timeNow).Scan(&todaysQuestion)
	if err != nil {
		return nil, err
	}

	todaysInfo := &models.TodaysInfo{
		Question: todaysQuestion,
		Answers:  answers,
	}

	return todaysInfo, nil
}

const existAnswer = `
SELECT EXISTS
(SELECT * FROM users_questions 
WHERE Question_id = $1 AND User_nickname = $2)`

const insertUserQuestion = `
INSERT INTO users_questions (Question_id, User_nickname, Answer, Created_at, Updated_at) 
VALUES ($1, $2, $3, $4, $5)`

func (store *Store) CreateAnswer(ctx context.Context, data models.UserQuestion) error {
	var exists int
	timeNow := time.Now().Format("02.01.2006 15:04:05")

	if !data.Valid() {
		return errors.ErrDataNotValid
	}

	err := store.Db.QueryRowContext(ctx, existAnswer, data.QuestionId, data.UserNickname).Scan(&exists)
	if err != nil {
		return err
	}

	if exists == 1 {
		return err
	} else {
		_, err = store.Db.ExecContext(ctx, insertUserQuestion, data.QuestionId, data.UserNickname, data.Answer, timeNow, timeNow)
		if err != nil {
			return err
		}
	}

	return nil
}

const updateAnswer = `
UPDATE users_questions SET Answer = $1, Updated_at = $2
WHERE Question_id = $3 AND User_nickname = $4`

func (store *Store) UpdateAnswer(ctx context.Context, data models.UserQuestion) error {
	updatedAt := time.Now().Format("02.01.2006 15:04:05")
	row, err := store.Db.ExecContext(ctx, updateAnswer, data.Answer, updatedAt, data.QuestionId, data.UserNickname)
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

const deleteAnswer = `
DELETE FROM users_questions 
WHERE Question_Id = $1 AND User_nickname = $2`

func (store *Store) DeleteAnswer(ctx context.Context, data models.UserQuestion) error {
	_, err := store.Db.ExecContext(ctx, deleteAnswer, data.QuestionId, data.UserNickname)
	if err != nil {
		return err
	}

	return nil
}
