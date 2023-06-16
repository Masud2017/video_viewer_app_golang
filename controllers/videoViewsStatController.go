package controllers

import (
	"github.com/gin-gonic/gin"
	"viewer_app.com/packages/services"
)

type VideoViewsStatController struct {

}

var videoStat = new(services.VideoStat)

func (videoViewsStatController *VideoViewsStatController) GetYoutubeVideoStat(c *gin.Context) {
	url := c.Query("url")

	urlInfo := videoStat.GetYoutubeVideoStat(url)

	jsonData := &urlInfo 
	c.JSON(200, gin.H{
		"data":jsonData,
	})
}

func (videoViewsStatController *VideoViewsStatController) GetInstagramVideoStat(c *gin.Context) {
	url := c.Query("url")

	urlInfo := videoStat.GetInstagramVideoStat(url)

	jsonData := &urlInfo 
	c.JSON(200, gin.H{
		"data":jsonData,
	})
}


func (videoViewsStatController *VideoViewsStatController) GetTiktokVideoStat(c *gin.Context) {
	url := c.Query("url")

	urlInfo := videoStat.GetTiktokVideoStat(url)

	jsonData := &urlInfo 
	c.JSON(200, gin.H{
		"data":jsonData,
	})
}