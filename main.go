package main

import (
	"github.com/gin-gonic/gin"
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
	server := gin.Default()

	api := server.Group("/api/v1")
	{
		user := api.Group("/auth")
		{
			user.POST("/login")
		}
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run()
}
