package main

import (
	"context"
	"github.com/nikolajjsj/go-jokes-api/api/config"
	"github.com/nikolajjsj/go-jokes-api/api/routes"
	"os"
)

func init() {
	// Initialize DB from Database.go under config folder
	Config.InitializeDB()
}

func main() {
	defer Config.Connection.Close(context.Background())

	// Fetch the port from the .env file
	port := os.Getenv("PORT")

	// Initialize router
	r := Routes.SetupRouter()
	r.Run(":" + port)
}
