package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
)

type CustomerRepoPort interface {
	GetCustomerByEmail(string) (queries.Customer, error)
	CreateCustomer(queries.CreateCustomerParams) (queries.Customer, error)
}

type CustomerHandlerPort interface {
	CreateCustomer(c *fiber.Ctx) error
	GetAllCustomers(c *fiber.Ctx) error
	GetCustomerByID(c *fiber.Ctx) error
	DeleteCustomer(c *fiber.Ctx) error
}

type CustomerApiPort interface {
	CreateCustomer(domain.Customer) (domain.Customer, error)
	GetAllCustomers(domain.CustomerParams) (domain.CustomersFetch, error)
	GetCustomerByID(string) (domain.Customer, error)
	DeleteCustomer(string) error
}
