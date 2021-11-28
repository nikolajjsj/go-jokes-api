package Controllers

import (
	"net/http"

	"api/Models"
	"github.com/gin-gonic/gin"
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
