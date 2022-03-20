package testutils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func LoadDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "C:/Users/User/Desktop/Разработка/go-workspace/src/question/database/db_test.db")
	if err != nil {
		log.Fatal("Cannot open DB")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB")
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	sql, err := ioutil.ReadFile("C:/Users/User/Desktop/Разработка/go-workspace/src/question/testutils/test_init.sql")
	if err != nil {
		log.Fatal("Cannot init DB")
	}

	_, err = db.Query(string(sql))
	if err != nil {
		log.Fatal("Cannot execute query")
	}

	return db
}

func ClearDatabase(s *sql.DB) {
	err := s.QueryRow(
		`DROP TABLE questions
		DROP TABLE users_auth
		DROP TABLE users_data
		DROP TABLE users_questions`)
	if err != nil {
		log.Fatal("Cannot drop database")
	}
}
