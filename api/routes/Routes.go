package Routes

import (
	"api/Controllers"
	"github.com/gin-gonic/gin"
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
