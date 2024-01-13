package main

import (
	"github.com/tharaka911/go-crud-1/initializers"
	"github.com/tharaka911/go-crud-1/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}