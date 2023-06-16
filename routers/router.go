package routers

import (
	"github.com/gin-gonic/gin"
	"viewer_app.com/packages/controllers"
)

type Router struct {}

func (router Router) GetRouter()  *gin.Engine {
	router := gin.Default();

	router.GET("/ping", controllers.Index)

	return router;
}