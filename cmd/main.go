package main

import (
	"log"
	"mephiSRW"
	"mephiSRW/pkg/handler"
	"mephiSRW/pkg/repository"
	"mephiSRW/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(mephiSRW.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Problem with start server, because %s", err.Error())
	}
}
