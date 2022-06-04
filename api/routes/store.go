package routes

import (
	"database/sql"
	"io/ioutil"

	"github.com/alekstet/question/conf"
	"github.com/alekstet/question/database"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

func New(db *sql.DB, store database.Querier) *Store {
	return &Store{
		Db:      db,
		Log:     logrus.New(),
		Routes:  httprouter.New(),
		Querier: store,
	}
}

type Store struct {
	Db      *sql.DB
	Log     *logrus.Logger
	Routes  *httprouter.Router
	Querier database.Querier
}

func InitDB(c *conf.ConfigDatabase) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", c.DbPath)
	if err != nil {
		return nil, err
	}

	sqlScript, err := ioutil.ReadFile(c.SQLInitPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Query(string(sqlScript))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
