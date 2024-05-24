package server

import (
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	//Cors Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//CSRF Middleware
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-CSRF-Token",
		CookieName:     "csrf_",
		CookieSecure:   true,
		CookieHTTPOnly: true,
	}))

	//Helmet
	app.Use(helmet.New())

	//Logger Middleware
	app.Use(logger.New())

	//Compress Middleware
	app.Use(compress.New())

	// Apply Limiter middleware
	app.Use(limiter.New(limiter.Config{
		// Max number of requests per duration
		Max: 30,
		// Duration for the above limit
		Expiration: 30 * time.Second,
		// Key to distinguish between different clients
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		// Handler to execute when limit is reached
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests, please try again later.",
			})
		},
	}))

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
