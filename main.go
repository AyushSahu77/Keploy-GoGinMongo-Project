package main

import (
	//default go packages
	"context"
	"time"

	//"net/http"

	// the project packages
	"example.com/ayush-keploy-apis/controllers"
	"example.com/ayush-keploy-apis/services"

	"github.com/gin-gonic/gin"

	//keploy packages

	"github.com/keploy/go-sdk/integrations/kgin/v1"
	//"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/keploy/go-sdk/integrations/kmongo"
	"github.com/keploy/go-sdk/keploy"
	"go.uber.org/zap"

	//mongo packages
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server      *gin.Engine
	us          services.UserService
	uc          controllers.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
	col         *kmongo.Collection
	logger      *zap.Logger
)

func New(host, db string) (*mongo.Client, error) {
	clientOptions := options.Client()

	clientOptions.ApplyURI("mongodb://" + host + "/" + db + "?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, clientOptions)
}

func main() {
	// defer mongoclient.Disconnect(ctx)

	logger, _ = zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	dbName, collection := "KeployInteGO", "user"
	client, err := New("localhost:27017", dbName)
	if err != nil {
		logger.Fatal("failed to create mongo db client", zap.Error(err))
	}

	db := client.Database(dbName)
	col = kmongo.NewCollection(db.Collection(collection))

	us = services.NewUserService(col)
	uc = controllers.New(us)

	port := "9090"

	k := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "ayush-keploy-apis",
			Port: port,
		},
		Server: keploy.ServerConfig{
			URL: "http://localhost:6789/api",
		},
	})

	r := gin.Default()

	kgin.GinV1(k, r)

	r.POST("/crud/user/create", uc.CreateUser)
	r.GET("/crud/user/get/:name", uc.GetUser)
	r.GET("/crud/user/getall", uc.GetAll)
	r.PATCH("/crud/user/update", uc.UpdateUser)
	r.DELETE("/crud/user/delete/:name", uc.DeleteUser)

	r.Run(":" + port)
}
