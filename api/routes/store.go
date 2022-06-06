package routes

import (
	"database/sql"
	"io/ioutil"

	"github.com/alekstet/question/conf"
	"github.com/alekstet/question/database"
	"github.com/alekstet/question/token"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db         *sql.DB
	Querier    database.Querier
	Log        *logrus.Logger
	Routes     *httprouter.Router
	TokenMaker token.TokenMaker
}

func New(db *sql.DB, store database.Querier, key string) *Store {
	jwt := &token.JWTTokenMaker{
		SymmetricKey: key,
	}

	return &Store{
		Db:         db,
		Querier:    store,
		Log:        logrus.New(),
		Routes:     httprouter.New(),
		TokenMaker: jwt,
	}
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
