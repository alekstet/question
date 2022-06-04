package main

import (
	"log"

	"github.com/alekstet/question/cmd"
	"github.com/alekstet/question/conf"
)

func main() {
	cnf, err := conf.Cnf()
	if err != nil {
		log.Fatalf("error with config: %s", err)
	}

	cmd.Run(cnf)
}
