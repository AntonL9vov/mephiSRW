package main

import (
	"github.com/spf13/viper"
	"log"
	"mephiSRW"
	"mephiSRW/pkg/handler"
	"mephiSRW/pkg/repository"
	"mephiSRW/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error intializing config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(mephiSRW.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Problem with start server, because %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
