package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajjsj/go-jokes-api/api/controllers"
	"net/http"
)

// SetupRouter function for initiating the Gin router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Homepage of the api
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to go-jokes-api!\nPlease checkout \"/jokes\"")
	})

	// Router endpoints
	router.GET("/jokes", Controllers.GetJokes)
	router.GET("/jokes/random", Controllers.RandomJoke)
	router.GET("/jokes/:id", Controllers.GetJokeByID)

	// Return GIN router
	return router
}
