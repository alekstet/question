package testutils

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

func LoadDatabase() *sql.DB {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../database/db_test.db")

	db, err := sql.Open("sqlite3", dir)
	if err != nil {
		log.Fatal("Cannot open DB")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB")
	}

	sql, err := ioutil.ReadFile("C:/Users/atete/go/src/question/testutils/test_init.sql")
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
	err := s.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB")
	}
	_, err = s.Query(`
	DELETE FROM users_auth;
	DELETE FROM users_auth;
	DELETE FROM users_data;
	DELETE FROM users_questions;
	`)
	if err != nil {
		log.Fatal("Cannot delete from database")
	}
}
