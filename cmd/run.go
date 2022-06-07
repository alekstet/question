package cmd

import (
	"log"

	"github.com/alekstet/question/api/routes"
	"github.com/alekstet/question/conf"
	"github.com/alekstet/question/database"
)

func Run(cnf *conf.ConfigDatabase) {
	config, err := conf.Cnf()
	if err != nil {
		log.Fatal(err)
	}

	db, err := routes.InitDB(config)
	if err != nil {
		log.Fatal(err)
	}

	dbStore := database.NewStore(db)
	store := routes.New(dbStore, config)
	routes.Routes(*store)

	store.Routes.Run(cnf.Host + ":" + cnf.Port)
}
