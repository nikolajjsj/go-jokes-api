package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajjsj/go-jokes-api/api/controllers"
)

// SetupRouter function for initiating the Gin router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Router endpoints
	router.GET("/jokes", Controllers.GetJokes)
	router.GET("/jokes/random", Controllers.RandomJoke)
	router.GET("/jokes/:id", Controllers.GetJokeByID)

	// Return GIN router
	return router
}
