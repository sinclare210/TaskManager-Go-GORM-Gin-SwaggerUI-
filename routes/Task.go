package routes

import (
	"TaskManager-Go-GORM-Gin-SwaggerUI/services"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewTask struct {
	Title   string     `json:"title"`
	Status  string     `json:"status"`
	DueDate *time.Time `json:"due_date"`
}

// CreateTask godoc
// @Summary      Create a new task
// @Description  Allows an authenticated user to create a task
// @Tags         Tasks
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
// @Tags         Tasks
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
// @Param        status   query     string  True "Task status filter (e.g. 'pending', 'complete')"
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

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Deletes a task by its ID, only if it belongs to the authenticated user
// @Tags         Tasks
// @Accept       json
// @Produce      json
// @Param        id   query     int  true  "Task ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      401  {object}  Response
// @Failure      403  {object}  Response
// @Failure      404  {object}  Response
// @Failure      500  {object}  Response
// @Security     TokenAuth
// @Router       /task/delete [delete]
func DeleteTask(context *gin.Context) {

	Id, err := strconv.ParseUint(context.Query("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, Response{
			Message: "Invalid task ID",
		})
		return
	}

	idVal, exists := context.Get("Id")
	if !exists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Message: "Unauthorized: user ID not found",
		})
		return
	}
	userID := idVal.(uint)

	task, err := services.GetTaskById(uint(Id))
	if err != nil {
		context.JSON(http.StatusNotFound, Response{
			Message: "Task not found",
		})
		return
	}

	if task.UserID != userID {
		context.JSON(http.StatusForbidden, Response{
			Message: "You are not allowed to delete this task",
		})
		return
	}

	if err := services.DeleteTask(uint(Id)); err != nil {
		context.JSON(http.StatusInternalServerError, Response{
			Message: "Failed to delete task",
		})
		return
	}

	context.JSON(http.StatusOK, Response{
		Message: "Task deleted successfully",
	})
}


// UpdateTask godoc
// @Summary Update a task
// @Description Update an existing task owned by the authenticated user
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param request body NewTask true "Update Task"
// @Success 200 {object} Response 
// @Failure 400 {object} Response 
// @Failure 401 {object} Response 
// @Failure 404 {object} Response 
// @Failure 500 {object} Response 
// @Security TokenAuth
// @Router /task/{id} [put]
func UpdateTask(context *gin.Context) {
	
	taskID, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, Response{Message: "Invalid task ID"})
		return
	}

	idVal, exists := context.Get("Id")
	if !exists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, Response{Message: "Unauthorized: user ID not found"})
		return
	}
	userID := idVal.(uint)

	
	var req NewTask
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, Response{Message: "Invalid request body"})
		return
	}

	
	err = services.UpdateTask(uint(taskID), userID, req.Title, req.Status, req.DueDate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, Response{Message: "Task not found"})
			return
		}
		context.JSON(http.StatusInternalServerError, Response{Message: "Failed to update task"})
		return
	}

	context.JSON(http.StatusOK, Response{Message: "Task updated successfully"})
}
