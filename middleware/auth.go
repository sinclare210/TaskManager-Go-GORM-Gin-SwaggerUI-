package middleware

import (
	"TaskManager-Go-GORM-Gin-SwaggerUI/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticated(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization token is required",
		})
		return
	}

	email, Id, err := auth.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or expired token",
		})
		return
	}
	context.Set("Email", email)
	log.Println(email)
	context.Set("Id", Id)
	log.Println(Id)

	context.Next()

}
