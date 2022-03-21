package testutils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
)

func LoadDatabase() *sql.DB {
	/* _, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../../database/db_test.db")
	if err := os.Chdir(dir); err != nil {
		log.Fatalf("Cannot change dir | %v | ", err)
	}

	fmt.Println(dir) */

	fmt.Println("Try opening DB")
	db, err := sql.Open("sqlite3", "C:/Users/atete/go/src/question/database/db_test.db")
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

	fmt.Println("DB OK")

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
