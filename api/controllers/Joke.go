package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-jokes-api/Models"
)

// Get Jokes ...
func GetJokes(c *gin.Context) {
	jokes, err := Models.GetAllJokes()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, jokes)
	}
}

// Create Joke ...

// Get Joke by id ...

// Update Jokes ...

// Delete joke ...
