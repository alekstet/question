package cmd

import (
	"github.com/alekstet/question/api/routes/process"
	"log"
	"net"
	"net/http"

	"github.com/alekstet/question/api/routes"
	"github.com/alekstet/question/conf"
	"github.com/alekstet/question/database"
)

func Run(cnf *conf.ConfigDatabase) {
	config, err := conf.Cnf()
	if err != nil {
		log.Fatal(err)
	}

	db, err := process.InitDB(config)
	if err != nil {
		log.Fatal(err)
	}

	dbStore := database.NewStore(db)

	store := process.New(db, dbStore)

	listener, err := net.Listen("tcp", cnf.Host+":"+cnf.Port)
	if err != nil {
		store.Log.Fatal(err)
	}

	routes.Routes(*store)

	server := &http.Server{
		Addr:    cnf.Host + cnf.Port,
		Handler: store.Routes,
	}
	store.Log.Info("Server is running on ", cnf.Host+":"+cnf.Port)
	store.Log.Fatal(server.Serve(listener))
}
