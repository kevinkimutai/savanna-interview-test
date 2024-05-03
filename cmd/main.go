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

	DBURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		POSTGRES_USERNAME,
		POSTGRES_PASSWORD,
		"localhost",
		DATABASE_PORT,
		"orderdb")

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
