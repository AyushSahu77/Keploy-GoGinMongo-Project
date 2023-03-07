package services

import (
	"context"
	"errors"

	"example.com/ayush-keploy-apis/models"
	"github.com/keploy/go-sdk/integrations/kmongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServiceImpl struct {
	usercollection *kmongo.Collection
}

func NewUserService(usercollection *kmongo.Collection) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
	}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	_, err := u.usercollection.InsertOne(ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(ctx context.Context, name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(users) == 0 {
		return nil, errors.New("no data present in the database found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(ctx context.Context, user *models.User) error {
	filter := bson.D{primitive.E{Key: "name", Value: user.Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.Name}, primitive.E{Key: "age", Value: user.Age}, primitive.E{Key: "address", Value: user.Address}}}}
	result, _ := u.usercollection.UpdateOne(ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no data can be found for the given name to update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, name *string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result, _ := u.usercollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no data can be found for the given name to delete")
	}
	return nil
}
