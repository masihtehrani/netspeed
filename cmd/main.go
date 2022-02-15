package main

import (
	"log"

	"github.com/masihtehrani/netspeed/server"
)

func main() {
	err := server.New()
	if err != nil {
		log.Panic(err)
	}
}
