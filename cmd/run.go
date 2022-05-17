package cmd

import (
	"net"
	"net/http"

	"github.com/alekstet/question/api/routes"
	"github.com/alekstet/question/conf"
)

func Run(cnf *conf.Config) {
	store := conf.New(cnf.SessionsKey)

	if err := store.InitDB(cnf); err != nil {
		store.Log.Fatal(err)
	}

	listener, err := net.Listen("tcp", cnf.Ip+":"+cnf.Port)
	if err != nil {
		store.Log.Fatal(err)
	}

	routes.Routes(*store)

	server := &http.Server{
		Addr:    cnf.Ip + cnf.Port,
		Handler: store.Routes,
	}
	store.Log.Info("Server is running on ", cnf.Ip+":"+cnf.Port)
	store.Log.Fatal(server.Serve(listener))
}
