package controllers

import ("github.com/gin-gonic/gin")


type AuthController struct {

}


func (authController AuthController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
