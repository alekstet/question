package main

import (
	"log"

	"question/cmd"
	"question/conf"
)

func main() {
	cnf, err := conf.Cnf()
	if err != nil {
		log.Fatalf("error with config: %s", err)
	}
	cmd.Run(cnf)
}
