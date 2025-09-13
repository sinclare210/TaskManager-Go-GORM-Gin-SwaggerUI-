package main

import (
	"TaskManager-Go-GORM-Gin-SwaggerUI/db"
	_ "TaskManager-Go-GORM-Gin-SwaggerUI/docs"
	"TaskManager-Go-GORM-Gin-SwaggerUI/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Task Manager
// @version         1.0
// @description     This is a note taker API .

// @contact.name   Sinclair
// @contact.url    https://x.com/syncc_crypt
// @contact.email  olajuwonsinclair@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey TokenAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	db.Init()
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load")
	}
	server := gin.Default()

	api := server.Group("/api/v1")
	{
		user := api.Group("/auth")

		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.LoginUser)
		}
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run()

}
