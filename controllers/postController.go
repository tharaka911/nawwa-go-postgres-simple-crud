package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tharaka911/go-crud-1/initializers"
	"github.com/tharaka911/go-crud-1/models"
)

//create post
func PostCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)

	//create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return post

	c.JSON(200, gin.H{
		"message": post,
	})
}

//get all posts
func PostGetAll(c *gin.Context) {

	//create posts list to store posts
	var posts []models.Post
	//find all posts
	initializers.DB.Find(&posts)

	//return posts
	c.JSON(200, gin.H{
		"message": posts,
	})
}

//get post by id
func PostGetById(c *gin.Context) {

	//create post variable to store post
	var post models.Post

	//get id from url
	id := c.Param("id")

	//find post by id
	initializers.DB.First(&post, id)

	//return post
	c.JSON(200, gin.H{
		"message": post,
	})
}

//update post
func PostUpdate(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"message": post,
	})
}

//delete post
func PostDelete(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{
		"message": "post deleted",
	})
}
