package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/kevinkimutai/savanna-app/internal/adapters/auth"
	"github.com/kevinkimutai/savanna-app/internal/adapters/db"
	handler "github.com/kevinkimutai/savanna-app/internal/adapters/handlers"
	"github.com/kevinkimutai/savanna-app/internal/adapters/rabbitmq"
	"github.com/kevinkimutai/savanna-app/internal/adapters/server"
	application "github.com/kevinkimutai/savanna-app/internal/app/core/api"
)

func main() {
	// Init Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	//Env Variables
	APP_PORT := os.Getenv("APP_PORT")
	POSTGRES_USERNAME := os.Getenv("POSTGRES_USERNAME")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DATABASE_PORT := os.Getenv("DB_PORT")
	DATABASE_HOST := os.Getenv("DB_HOST")
	DATABASE_NAME := os.Getenv("DB_NAME")
	DATABASE_TEST := os.Getenv("TEST_DB")

	var DB string

	if os.Getenv("ENV") != "development" {
		DB = DATABASE_NAME
	} else {
		DB = DATABASE_TEST
	}

	fmt.Println(DB)
	DBURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		POSTGRES_USERNAME,
		POSTGRES_PASSWORD,
		DATABASE_HOST,
		DATABASE_PORT,
		DB)

	RABBITMQSERVER := os.Getenv("RABBITMQ_SERVER")

	//Dependency injection

	//Database
	//Connect To DB
	dbAdapter := db.NewDB(DBURL)

	//RabbitMQ
	msgQueue := rabbitmq.NewRabbitMQServer(RABBITMQSERVER)

	//Repositories
	//customerRepo := application.NewCustomerRepo(dbAdapter)
	orderRepo := application.NewOrderRepo(dbAdapter, msgQueue)
	productRepo := application.NewProductRepo(dbAdapter)

	//Services
	//customerService := handler.NewCustomerService(customerRepo)
	orderService := handler.NewOrderService(orderRepo)
	productService := handler.NewProductService(productRepo)

	authService, err := auth.New(dbAdapter)
	if err != nil {
		log.Fatal(err)
	}

	//Server
	server := server.New(
		APP_PORT,
		authService,
		//customerService,
		orderService,
		productService)

	server.Run()

}
