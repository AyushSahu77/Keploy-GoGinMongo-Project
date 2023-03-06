package main

import (
	//default go packages
	"context"
	"fmt"
	"log"
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
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb+srv://AyushsCluster77:mongocluster139@cluster0.wy6ry18.mongodb.net/GoApp?retryWrites=true&w=majority")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("GoAppDB").Collection("user")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	server = gin.Default()
}

func New(host, db string) (*mongo.Client, error) {
	clientOptions := options.Client()

	clientOptions.ApplyURI("mongodb://" + host + "/" + db + "?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, clientOptions)
}

func main() {
	// defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/crud")
	uc.RegisterUserRoutes(basepath)

	logger, _ = zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	dbName, collection := "KeployInteGO", "user"
	client, err := New("mongodb+srv%3a%2f%2fAyushsCluster77%3amongocluster139@cluster0.wy6ry18.mongodb.net%2fGoApp%3fretryWrites=true&w=majority", dbName)
	if err != nil {
		logger.Fatal("failed to create mongo db client", zap.Error(err))
	}

	db := client.Database(dbName)
	col = kmongo.NewCollection(db.Collection(collection))

	port := "9090"
	r := gin.New()
	k := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "ayush-keploy-apis",
			Port: port,
		},
		Server: keploy.ServerConfig{
			URL: "http://localhost:6789/api",
		},
	})
	kgin.GinV1(k, r)
	r.POST("/crud/user/create", uc.CreateUser)
	r.GET("/crud/user/get/:name", uc.GetUser)
	r.GET("/crud/user/getall", uc.GetAll)
	r.PATCH("/crud/user/update", uc.UpdateUser)
	r.DELETE("/crud/user/delete/:name", uc.DeleteUser)
	log.Fatal(r.Run(":" + port))
}
