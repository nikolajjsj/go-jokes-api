package main

import (
	"context"
	"github.com/nikolajjsj/go-jokes-api/api/config"
	"github.com/nikolajjsj/go-jokes-api/api/routes"
	"os"
)

func main() {
	// Initialize DB from Database.go under config folder
	Config.InitializeDB()
	defer Config.Connection.Close(context.Background())

	// Fetch the port from the .env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize router
	r := Routes.SetupRouter()
	r.Run(":" + port)
}
