package routes

import (
	"database/sql"
	"io/ioutil"

	"github.com/alekstet/question/conf"
	"github.com/alekstet/question/database"
	"github.com/alekstet/question/token"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Querier    database.Querier
	Log        *logrus.Logger
	Routes     *gin.Engine
	TokenMaker token.TokenMaker
	Config     conf.ConfigDatabase
}

func New(store database.Querier, config *conf.ConfigDatabase) *Store {
	jwt := &token.JWTTokenMaker{
		SymmetricKey: config.SymmetricKey,
	}

	return &Store{
		Querier:    store,
		Log:        logrus.New(),
		Routes:     gin.Default(),
		TokenMaker: jwt,
		Config:     *config,
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
