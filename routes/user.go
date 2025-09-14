package routes

import (
	"TaskManager-Go-GORM-Gin-SwaggerUI/auth"
	"TaskManager-Go-GORM-Gin-SwaggerUI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterUserF struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserF struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"Message"`
}

// Register User For the Task Manager
// @Summary Register User For the Task Manager
// @Description Register User For the Task Manager
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterUserF true "User Registration Data"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /auth/register [post]
func RegisterUser(context *gin.Context) {
	var user RegisterUserF

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, Response{Message: "Invalid Inputs"})
		return
	}

	err = services.RegisterUser(user.Name, user.Email, user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, Response{Message: "Database error"})
		return
	}

	context.JSON(http.StatusCreated, Response{Message: "Welcome, " + user.Name})
}

// Login a Registered User
// @Summary  Login a Registered User
// @Description  Login a Registered User
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginUserF true "User Registration Data"
// @Success 202 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /auth/login [post]
func LoginUser(context *gin.Context) {
	var user LoginUserF

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, Response{Message: "Invalid Inputs"})
		return
	}

	userInfo, err := services.LoginUser(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, Response{Message: "Database error"})
		return
	}

	err = auth.VerifyHash(userInfo.Password, user.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, Response{Message: "Invalid Credential"})
		return
	}

	token, err := auth.GenerateToken(userInfo.Email, userInfo.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, Response{Message: "Failedd to generate token"})
		return
	}

	context.JSON(http.StatusAccepted, Response{Message: "Welcome, " + userInfo.Name + " " + token})
}
