package main

import (
	//default go packages
	"context"
	"time"

	// the project packages
	"example.com/ayush-keploy-apis/controllers"
	"example.com/ayush-keploy-apis/services"

	"github.com/gin-gonic/gin"

	//keploy packages

	"github.com/keploy/go-sdk/integrations/kgin/v1"
	"github.com/keploy/go-sdk/integrations/kmongo"
	"github.com/keploy/go-sdk/keploy"
	"go.uber.org/zap"

	//mongo packages
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	us          services.UserService
	uc          controllers.UserController
	col         *kmongo.Collection
	logger      *zap.Logger
)

func New(host, db string) (*mongo.Client, error) {
	clientOptions := options.Client()

	clientOptions.ApplyURI("mongodb+srv://" + host + "/" + db + "?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, clientOptions)
}

func main() {
	// defer mongoclient.Disconnect(ctx)

	logger, _ = zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	dbName, collection := "GoAppDB", "user"       // Change this to your DB and Collection names
	client, err := New("AyushsCluster77:<PASSWORD>@cluster0.wy6ry18.mongodb.net", dbName)  // Change the <PASSWORD> to your password of the cluster
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

	r.POST("/create", uc.CreateUser)
	r.GET("/get/:name", uc.GetUser)
	r.GET("/getall", uc.GetAll)
	r.PATCH("/update", uc.UpdateUser)
	r.DELETE("/delete/:name", uc.DeleteUser)

	r.Run(":" + port)
}
