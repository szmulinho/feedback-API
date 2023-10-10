package main

import (
	"fmt"
	"github.com/szmulinho/feedback/internal/database"
	"github.com/szmulinho/feedback/internal/server"
	"github.com/szmulinho/feedback/internal/utils"
	"log"
)

func main() {
	fmt.Println("Starting the application...")
	defer fmt.Println("Closing the application")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	ctx, _, wait := utils.Gracefully()

	server.Run(ctx, db)

	wait()
}
