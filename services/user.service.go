package services

import (
	"example.com/ayush-keploy-apis/models"
	"golang.org/x/net/context"
)

type UserService interface {
	CreateUser(context.Context, *models.User) error
	GetUser(context.Context, *string) (*models.User, error)
	GetAll(context.Context) ([]*models.User, error)
	UpdateUser(context.Context, *models.User) error
	DeleteUser(context.Context, *string) error
}
