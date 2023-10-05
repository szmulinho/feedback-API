package main

import (
	"github.com/szmulinho/feedback/cmd/server"
	"github.com/szmulinho/feedback/internal/database"
)

func main() {

	database.Connect()

	server.Run()
}
