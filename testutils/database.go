package testutils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

func LoadDatabase() *sql.DB {
	_, filename, _, _ := runtime.Caller(0)
	path_base := path.Join(path.Dir(filename), "../database/db_test.db")
	path_init_db := path.Join(path.Dir(filename), "../conf/init.sql")

	db, err := sql.Open("sqlite3", path_base)
	if err != nil {
		log.Fatal("Cannot open DB")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Cannot ping DB")
	}

	query, err := ioutil.ReadFile(path_init_db)
	if err != nil {
		log.Fatal("Cannot init DB")
	}

	_, err = db.Query(string(query))
	if err != nil {
		log.Fatal("Cannot execute query")
	}

	return db
}

func ClearDatabase(db *sql.DB) {
	store := NewStore()
	store.Db = db
	if err := store.Db.Ping(); err != nil {
		log.Fatal("Cannot ping DB")
	}

	result, err := store.Db.Exec(`DELETE FROM users_auth;DELETE FROM users_data;DELETE FROM questions;DELETE FROM users_questions`)
	if err != nil {
		log.Fatal("Cannot delete from database")
	}

	fmt.Println(result.RowsAffected())
}
