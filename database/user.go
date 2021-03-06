package database

import (
	"fmt"
	"github.com/alekstet/question/api/models"
)

func (s Store) GetUsers() ([]models.UsersData, error) {
	fmt.Println("DB")
	rows, err := s.Db.Query("SELECT User_nickname, Name, Sex FROM users_data")
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

func (s *Store) GetUserInfo(nickname, sort string) (*models.UserInfo, error) {
	var name string
	var sex string
	usersAnsw := []models.UserAnsw{}
	sql := fmt.Sprintf(
		`SELECT Question_id, Question, Answer FROM users_questions 
		INNER JOIN questions ON Question_id = Date WHERE User_nickname = '%s'`, nickname)
	if sort == "dateup" {
		sql = fmt.Sprintf("%s ORDER BY Date ASC", sql)
	}

	if sort == "datedown" {
		sql = fmt.Sprintf("%s ORDER BY Date DESC", sql)
	}

	rows, err := s.Db.Query(sql)
	if err != nil {
		return &models.UserInfo{}, err
	}

	defer rows.Close()

	for rows.Next() {
		userAnsw := models.UserAnsw{}
		err := rows.Scan(&userAnsw.Date, &userAnsw.Question, &userAnsw.Answer)
		if err != nil {
			return &models.UserInfo{}, err
		}

		usersAnsw = append(usersAnsw, userAnsw)
	}

	err = s.Db.QueryRow("SELECT Name, Sex FROM users_data WHERE User_nickname = $1", nickname).Scan(&name, &sex)
	if err != nil {
		return &models.UserInfo{}, err
	}

	return &models.UserInfo{
		Name: name,
		Sex: sex,
		Answers: usersAnsw}, nil
}

