package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tharaka911/go-crud-1/controllers"
	"github.com/tharaka911/go-crud-1/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostCreate)
	r.GET("/posts", controllers.PostGetAll)
	r.GET("/posts/:id", controllers.PostGetById)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
