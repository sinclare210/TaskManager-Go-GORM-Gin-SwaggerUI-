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

// GetTask godoc
// @Summary      Get all tasks for the authenticated user
// @Description  Returns a list of tasks belonging to the logged-in user
// @Tags         tasks
// @Produce      json
// @Success      200  {array}   model.Task
// @Failure      401  {object}  Response
// @Failure      500  {object}  Response
// @Security TokenAuth
// @Router  /task/alltask [get]
func GetAllTask(context *gin.Context) {
	idVal, exists := context.Get("Id")
	if !exists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Message: "Unauthorized: user ID not found",
		})
		return
	}

	userID := idVal.(uint)

	tasks, err := services.GetTask(userID)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Message: "Failed to fetch tasks",
		})
		return
	}

	context.JSON(http.StatusOK, tasks)
}

// GetTask godoc
// @Summary      Get tasks for the authenticated user
// @Description  Retrieves all tasks for the logged-in user. You can filter by status using a query parameter.
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        status   query     string  True "Task status filter (e.g. 'pending', 'completed')"
// @Success      200      {array} model.Task
// @Failure      400      {object}  Response
// @Failure      401      {object}  Response
// @Failure      500      {object}  Response
// @Security TokenAuth
// @Router /task/specifictask [get]
func GetTask(context *gin.Context) {

	idVal, exists := context.Get("Id")
	if !exists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Message: "Unauthorized: user ID not found",
		})
		return
	}
	userID := idVal.(uint)

	status := context.Query("status")

	tasks, err := services.GetTaskBasedOnStat(userID, status)
	if err != nil {
		context.JSON(http.StatusInternalServerError, Response{
			Message: "Failed to fetch tasks",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Tasks retrieved successfully",
		"tasks":   tasks,
	})
}
