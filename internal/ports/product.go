package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
)

type ProductRepoPort interface {
	CreateProduct(domain.Product) (domain.Product, error)
	GetAllProducts(queries.ListProductsParams) (domain.ProductsFetch, error)
	GetProduct(string) (domain.Product, error)
	DeleteProduct(string) error
}

type ProductHandlerPort interface {
	CreateProduct(c *fiber.Ctx) error
	GetAllProducts(c *fiber.Ctx) error
	GetProductByID(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}

type ProductApiPort interface {
	CreateProduct(domain.Product) (domain.Product, error)
	GetAllProducts(domain.ProductParams) (domain.ProductsFetch, error)
	GetProduct(string) (domain.Product, error)
	DeleteProduct(string) error
}
