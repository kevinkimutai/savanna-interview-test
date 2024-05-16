package server

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/kevinkimutai/savanna-app/docs"
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

// @title           Order API
// @version         1.0
// @description     Order API docs.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func (s *ServerAdapter) Run() {
	//Initialize Fiber
	app := fiber.New()

	//Logger Middleware
	app.Use(logger.New())

	//Swagger Middleware
	cfg := swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger Order API Docs",
	}
	app.Use(swagger.New(cfg))

	// Define routes
	//app.Route("/api/v1/customer", s.CustomerRouter)
	app.Route("/api/v1/order", s.OrderRouter)
	app.Route("/api/v1/product", s.ProductRouter)

	app.Listen(":" + s.port)
}
