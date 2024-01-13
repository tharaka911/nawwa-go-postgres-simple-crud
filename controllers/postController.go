package controllers

import "github.com/gin-gonic/gin"

func PostCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "nawwa this is great",
	})
}