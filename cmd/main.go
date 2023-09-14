package main

import (
	jwt "github.com/ShatALex/TestTaskBackDev"
	"github.com/ShatALex/TestTaskBackDev/pkg/handler"
	"github.com/ShatALex/TestTaskBackDev/pkg/repository"
	"github.com/ShatALex/TestTaskBackDev/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables:%s", err.Error())
	}

	db := repository.NewMongoDB()

	rep := repository.NewRepository(db, viper.GetString("mongo.user_collection"))
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	server := new(jwt.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		logrus.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
