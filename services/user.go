package services

import (
	"TaskManager-Go-GORM-Gin-SwaggerUI/auth"
	"TaskManager-Go-GORM-Gin-SwaggerUI/db"
	"TaskManager-Go-GORM-Gin-SwaggerUI/model"
)

func RegisterUser(name, email, password string) error {
	HashedPassword, err := auth.GenerateHash(password)
	if err != nil {
		return err
	}
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: HashedPassword,
	}
	return db.DB.Create(user).Error
}

func LoginUser(email string) (model.User, error) {
	var user model.User
	result := db.DB.Where("email = ?", email).First(&user)

	return user, result.Error

}
