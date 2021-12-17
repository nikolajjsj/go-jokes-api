package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajjsj/go-jokes-api/api/controllers"
	"net/http"
)

// SetupRouter function for initiating the Gin router
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// Gather templates + stylesheets
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	// Homepage of the api
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// Router endpoints
	router.GET("/jokes", Controllers.GetJokes)
	router.GET("/random", Controllers.RandomJoke)
	router.GET("/jokes/:id", Controllers.GetJokeByID)

	// Return GIN router
	return router
}
