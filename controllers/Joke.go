package Controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nikolajjsj/go-jokes-api/api/models"
)

// GetJokes get all jokes controller
func GetJokes(c *gin.Context) {
	jokes, err := Models.GetAllJokes()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, jokes)
	}
}

// Create Joke ...

// GetJokeByID get joke by its id controller
func GetJokeByID(c *gin.Context) {
	IDParam := c.Param("id")
	ID, err := strconv.Atoi(IDParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	joke, err := Models.GetJokeByID(ID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, joke)
	}
}

// RandomJoke controller for getting a random joke
func RandomJoke(c *gin.Context) {
	joke, err := Models.GetRandomJoke()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, joke)
	}
}
