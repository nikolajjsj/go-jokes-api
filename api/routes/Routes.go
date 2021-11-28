package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolajjsj/golang-jokes-api/Controllers"
)

// Setup router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/jokes")
	{
		grp1.GET("jokes", Controllers.GetJokes)
		// grp1.POST("jokes", Controllers.CreateJoke)
		// grp1.GET("jokes/:id", Controllers.GetJokeByID)
		// grp1.PUT("jokes/:id", Controllers.UpdateJoke)
		// grp1.DELETE("jokes/:id", Controllers.DeleteJoke)
	}
	return r
}
