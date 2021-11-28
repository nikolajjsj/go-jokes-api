package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikolajjsj/golang-jokes-api/Models"
)

// Get Jokes
func GetJokes(c *gin.Context) {
	var jokes []Models.Joke
	err := Models.GetAllJokes(&jokes)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, jokes)
	}
}

// Create Joke

// Get Joke by id

// Update Jokes

// Delete joke
