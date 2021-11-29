package main

import (
	"Abinash393/digital-trons-test/controller"
	"Abinash393/digital-trons-test/middleware"
	"Abinash393/digital-trons-test/model"

	"github.com/gin-gonic/gin"
)

func init() {
	model.InitializeDB()
	model.DB.AutoMigrate(&model.User{}, &model.Folder{}, &model.File{})
	model.Seed()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		api.POST("/users/signup", controller.Signup)
		api.POST("/users/login", controller.Login)
	}

	fs := api.Group("/fs")
	fs.Use(middleware.IsAuthorized)
	{

	}

	router.NoRoute(controller.NotFound)

	router.Run() // port#8080
}
