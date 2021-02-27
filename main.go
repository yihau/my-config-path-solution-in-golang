package main

import (
	"log"

	"github.com/my-config-path-solution-in-golang/config"
)

func main() {
	err := config.LoadConfig("./config.yml")
	if err != nil {
		log.Fatal(err)
	}
}
