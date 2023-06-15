package routers

import (
	"github.com/gin-gonic/gin"
	"viewer_app.com/packages/controllers"
)

func SetupRouter()  *gin.Engine {
	router := gin.Default();

	router.GET("/ping", controllers.Index)

	return router;
}