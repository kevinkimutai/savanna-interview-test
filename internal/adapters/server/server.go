package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kevinkimutai/savanna-app/internal/ports"
)

type ServerAdapter struct {
	port string
	auth ports.AuthHandlerPort
	//customer ports.CustomerHandlerPort
	order   ports.OrderHandlerPort
	product ports.ProductHandlerPort
}

func New(
	port string,
	auth ports.AuthHandlerPort,
	//customer ports.CustomerHandlerPort,
	order ports.OrderHandlerPort,
	product ports.ProductHandlerPort) *ServerAdapter {
	return &ServerAdapter{
		port: port,
		auth: auth,
		//customer: customer,
		order:   order,
		product: product}
}

func (s *ServerAdapter) Run() {
	//Initialize Fiber
	app := fiber.New()

	//Logger Middleware
	app.Use(logger.New())

	//Swagger Middleware
	// cfg := swagger.Config{
	// 	BasePath: "/api/v1/",
	// 	FilePath: "./docs/swagger/swagger.json",
	// 	Path:     "swagger",
	// 	Title:    "Swagger Movie API Docs",
	// }
	// app.Use(swagger.New(cfg))

	// Define routes
	//app.Route("/api/v1/customer", s.CustomerRouter)
	app.Route("/api/v1/order", s.OrderRouter)
	app.Route("/api/v1/product", s.ProductRouter)

	app.Listen(":" + s.port)
}
