package main

import (
	"context"
	"github.com/nikolajjsj/go-jokes-api/api/config"
	"github.com/nikolajjsj/go-jokes-api/api/routes"
)

func init() {
}

func main() {
	// Initialize DB from Database.go under config folder
	Config.InitializeDB()
	defer Config.Connection.Close(context.Background())

	// Initialize router
	r := Routes.SetupRouter()
	r.Run()
}
