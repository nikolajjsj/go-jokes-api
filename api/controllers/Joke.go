package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikolajjsj/go-jokes-api/api/models"
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
