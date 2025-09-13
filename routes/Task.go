package routes

import (
	"TaskManager-Go-GORM-Gin-SwaggerUI/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type NewTask struct {
	Title   string     `json:"title"`
	Status  string     `json:"status"`
	DueDate *time.Time `json:"due_date"`
}

// CreateTask godoc
// @Summary      Create a new task
// @Description  Allows an authenticated user to create a task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param  request body  NewTask  true  "Task to create"
// @Success      201   {object}  model.Task
// @Failure      400   {object}   Response
// @Failure      401   {object}   Response
// @Failure      500   {object}   Response
// @Security TokenAuth
// @Router  /task/create [post]
func CreateTask(context *gin.Context) {
	var newTask NewTask
	if err := context.ShouldBindJSON(&newTask); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Message: "Invalid request body",
		})
		return
	}

	userID, exists := context.Get("Id")
	if !exists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Message: "Unauthorized: user ID not found",
		})
		return
	}

	if err := services.CreateTask(userID.(uint), newTask.Title, newTask.Status, newTask.DueDate); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Message: "Failed to create task",
		})
		return
	}

	context.JSON(http.StatusCreated, Response{
		Message: "Task created successfully",
	})
}
