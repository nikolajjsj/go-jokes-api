package Routes

import (
	"github.com/gin-gonic/gin"
	"go-jokes-api/Controllers"
)

// Setup router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// /Jokes path
	router.GET("/jokes", Controllers.GetJokes)
	// router.POST("/jokes", Controllers.CreateJoke)
	// router.GET("/jokes/:id", Controllers.GetJokeByID)
	// router.PUT("/jokes/:id", Controllers.UpdateJoke)
	// router.DELETE("/jokes/:id", Controllers.DeleteJoke)

	// Return GIN router
	return router
}
