package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
)

type OrderRepoPort interface {
	CreateOrderItem(orderItem domain.OrderItem) (domain.OrderItem, error)
	CreateOrder(string, string) (domain.Order, error)
	GetTotalPrice(string) (float64, error)
	UpdateOrderTotalPrice(string, float64) (domain.Order, error)
	GetOrderByID(orderID string) (domain.Order, error)
	DeleteOrder(orderID string) error
	GetAllOrders(params queries.ListOrdersParams) (domain.OrdersFetch, error)
}

type OrderHandlerPort interface {
	CreateOrder(c *fiber.Ctx) error
	GetAllOrders(c *fiber.Ctx) error
	GetOrderByID(c *fiber.Ctx) error
	DeleteOrder(c *fiber.Ctx) error
}

type OrderApiPort interface {
	CreateOrder([]domain.OrderItem, string) (domain.Order, error)
	GetAllOrders(domain.OrderParams) (domain.OrdersFetch, error)
	GetOrderByID(string) (domain.Order, error)
	DeleteOrder(string) error
}
