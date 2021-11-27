package main

import (
	"fmt"
	"github.com/nikolajjsj/golang-jokes-api/Config"
	"github.com/nikolajjsj/golang-jokes-api/Models"
	"github.com/nikolajjsj/golang-jokes-api/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Joke{})

	r := Routes.SetupRouter()
	r.Run()
}
