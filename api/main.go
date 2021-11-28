package main

import (
	"context"
	"github.com/nikolajjsj/golang-jokes-api/Config"
	"github.com/nikolajjsj/golang-jokes-api/routes"
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
