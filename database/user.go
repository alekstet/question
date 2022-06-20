package database

import (
	"context"
	"fmt"

	"github.com/alekstet/question/api/models"
)

func (store *Store) GetUsers(ctx context.Context) ([]models.UsersData, error) {
	rows, err := store.Db.QueryContext(ctx, "SELECT User_nickname, Name, Sex FROM users_data")
	if err != nil {
		return []models.UsersData{}, err
	}

	defer rows.Close()

	users := []models.UsersData{}

	for rows.Next() {
		usersData := models.UsersData{}
		err := rows.Scan(&usersData.UserNickname, &usersData.Name, &usersData.Sex)
		if err != nil {
			return []models.UsersData{}, err
		}

		users = append(users, usersData)
	}

	return users, nil
}

const (
	queryGetUserInfo = `
SELECT Question_id, Question, Answer 
FROM users_questions 
INNER JOIN questions ON Question_id = Date 
WHERE User_nickname = '%s'`

	queryGetByUsername = `
SELECT Name, Sex FROM users_data 
WHERE User_nickname = $1`
)

func (store *Store) GetUserInfo(ctx context.Context, nickname, sort string) (*models.UserInfo, error) {
	var name, sex string
	usersAnsw := []models.UserAnsw{}
	sql := fmt.Sprintf(queryGetUserInfo, nickname)
	if sort == "dateup" {
		sql = fmt.Sprintf("%s ORDER BY Date ASC", sql)
	}

	if sort == "datedown" {
		sql = fmt.Sprintf("%s ORDER BY Date DESC", sql)
	}

	rows, err := store.Db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		userAnsw := models.UserAnsw{}
		err := rows.Scan(&userAnsw.Date, &userAnsw.Question, &userAnsw.Answer)
		if err != nil {
			return nil, err
		}

		usersAnsw = append(usersAnsw, userAnsw)
	}

	err = store.Db.QueryRowContext(ctx, queryGetByUsername, nickname).Scan(&name, &sex)
	if err != nil {
		return nil, err
	}

	return &models.UserInfo{
		Name:    name,
		Sex:     sex,
		Answers: usersAnsw}, nil
}
