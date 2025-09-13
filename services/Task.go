package services

import (
	"TaskManager-Go-GORM-Gin-SwaggerUI/db"
	"TaskManager-Go-GORM-Gin-SwaggerUI/model"
	"fmt"
	"time"
)

func CreateTask(UserID uint, Title, Status string, DueDate *time.Time) error {
	newTask := &model.Task{
		UserID:  UserID,
		Title:   Title,
		Status:  Status,
		DueDate: DueDate,
	}

	err := db.DB.Create(newTask).Error
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func GetTask(userId uint) ([]model.Task, error) {
	var task []model.Task

	result := db.DB.Where("user_id", userId).Find(&task)

	return task, result.Error
}

func GetTaskBasedOnStat(userId uint, status string) ([]model.Task, error) {
	var tasks []model.Task

	result := db.DB.Where("user_id = ? AND status = ?", userId, status).Find(&tasks)

	return tasks, result.Error
}
