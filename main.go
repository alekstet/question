package main

import (
	"log"

	"github.com/question/cmd"
	"github.com/question/conf"
)

func main() {
	cnf, err := conf.Cnf()
	if err != nil {
		log.Fatalf("error with config: %s", err)
	}
	cmd.Run(cnf)
}
