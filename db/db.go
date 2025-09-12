package db

import (
	"TaskManager-Go-GORM-Gin-SwaggerUI/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("could not connect to database")
	}

	err = DB.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		panic("could not create table")
	}
}
