package main

import (
	"fmt"
	"log"

	"question/cmd"
	"question/conf"
)

func main() {
	cnf, err := conf.Cnf()
	if err != nil {
		fmt.Println("fatal")
		log.Fatalf("error with config: %s", err)
	}
	cmd.Run(cnf)
}
