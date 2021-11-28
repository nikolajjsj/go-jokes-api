package main

import (
	"api/Config"
	"api/routes"
	"context"
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
