package services

import "example.com/ayush-keploy-apis/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
