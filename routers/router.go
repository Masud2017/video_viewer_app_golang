package routers

import (
	"github.com/gin-gonic/gin"
	"viewer_app.com/packages/controllers"
)

func SetupRouter()  *gin.Engine {
	router := gin.Default();

	viewsStatController := new(controllers.VideoViewsStatController)

	router.GET("/youtubestat", viewsStatController.GetYoutubeVideoStat)
	router.GET("/instagramstat", viewsStatController.GetInstagramVideoStat)
	router.GET("/tiktokstat", viewsStatController.GetTiktokVideoStat)

	return router;
}