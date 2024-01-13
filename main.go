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

	r.GET("/post", controllers.PostCreate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
