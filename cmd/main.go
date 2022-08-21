package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/AsaHero/simple-bank/api"
	"github.com/AsaHero/simple-bank/api/handler"
	"github.com/AsaHero/simple-bank/config"
	"github.com/AsaHero/simple-bank/db/driver"
	"github.com/AsaHero/simple-bank/service"
	"github.com/AsaHero/simple-bank/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error on loading .env file: %s", err.Error())
	}
	config := config.Load()

	db, err := driver.InitPostgresDB(&config)
	if err != nil {
		log.Fatalf("error on database conntection: %s", err.Error())
	}
	storage := storage.NewStorage(db)
	service := service.NewService(storage)
	handler := handler.NewHandler(service)
	router := api.InitRouter(handler)

	server := api.Server{}

	if err = server.Start(config.Host+":"+config.Port, router); err != nil {
		log.Fatalf("error on starting the server: %s", err.Error())
	}

}
