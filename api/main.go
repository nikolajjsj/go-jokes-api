package main

import (
	"context"
	"go-jokes-api/Config"
	"go-jokes-api/routes"
)

func init() {
}

func main() {
	// Initialize DB from Database.go under Config folder
	Config.InitializeDB()
	defer Config.Connection.Close(context.Background())

	// Initialize router
	r := Routes.SetupRouter()
	r.Run()
}
