package cmd

import (
	"net"
	"net/http"
	"question/api/routes"
	"question/conf"
	"question/testutils"

	"github.com/gorilla/sessions"
)

func Run(cnf *conf.Config) {
	session := sessions.NewCookieStore([]byte(cnf.SessionsKey))
	store := conf.New(session)

	err := store.InitDB(cnf)
	if err != nil {
		store.Log.Fatal(err)
	}

	listener, err := net.Listen("tcp", cnf.Ip+":"+cnf.Port)
	if err != nil {
		store.Log.Fatal(err)
	}

	testutils.LoadDatabase()

	routes.Routes(*store)

	server := &http.Server{
		Addr:    cnf.Ip + cnf.Port,
		Handler: store.Routes,
	}
	store.Log.Info("Server is running on ", cnf.Ip+":"+cnf.Port)
	store.Log.Fatal(server.Serve(listener))
}
