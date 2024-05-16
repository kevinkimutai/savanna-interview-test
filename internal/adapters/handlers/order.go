package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/ports"
)

type OrderService struct {
	api ports.OrderApiPort
}

func NewOrderService(api ports.OrderApiPort) *OrderService {
	return &OrderService{api: api}
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Creates a new order with the provided order items.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      int    true  "Customer ID"
// @Param        orderItems body []domain.OrderItemRequest true "Order items to be added"
// @Success      201  {object} domain.OrderResponse "Returns the newly created order"
// @Failure      400  {object} domain.ErrorResponse      "Bad request, invalid input"
// @Failure      404  {object} domain.ErrorResponse      "Customer not found"
// @Failure      500  {object} domain.ErrorResponse      "Internal server error"
// @Router       /order [post]
func (s *OrderService) CreateOrder(c *fiber.Ctx) error {
	//Get CustomerID
	cus := c.Locals("customer")

	customer, ok := cus.(queries.Customer)
	if !ok {
		fmt.Println("Type assertion failed, cus is not of type queries.Customer")

	}

	//Receive OrderItems
	orderItems := []domain.OrderItemRequest{}

	//Bind To struct
	if err := c.BodyParser(&orderItems); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	// Generate uuid
	uuid := domain.GenerateUUID()

	//Add UUID to each struct
	items := addUUID(uuid, orderItems)

	//Api
	order, err := s.api.CreateOrder(items, customer.CustomerID)
	if err != nil {
		return c.Status(500).JSON(domain.ErrorResponse{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	return c.Status(201).JSON(
		domain.OrderResponse{
			StatusCode: 201,
			Message:    "Successfully created order",
			Data:       order})
}

// GetAllOrders godoc
// @Summary      Get all orders
// @Description  Retrieves all orders based on the provided query parameters.
// @Tags         orders
// @Produce      json
// @Param        page query int false "Page number"
// @Param        limit query int false "Number of orders per page"
// @Param        sortBy query string false "Sort by field"
// @Param        sortOrder query string false "Sort order ('asc' or 'desc')"
// @Success      200  {object} domain.OrdersResponse "Successfully retrieved orders"
// @Failure      500  {object} domain.ErrorResponse  "Internal server error"
// @Router       /order [get]
func (s *OrderService) GetAllOrders(c *fiber.Ctx) error {
	//Get Query Params
	m := c.Queries()

	//Bind To OrderParams
	orderParams := domain.CheckOrderParams(m)

	//Get All Orders API
	orders, err := s.api.GetAllOrders(orderParams)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}
	return c.Status(200).JSON(
		domain.OrdersResponse{
			StatusCode:    200,
			Message:       "Successfully retrieved order",
			Page:          orders.Page,
			NumberOfPages: orders.NumberOfPages,
			Total:         orders.Total,
			Data:          orders.Data,
		})
}

// GetOrderByID godoc
// @Summary      Get order by ID
// @Description  Retrieves an order based on the provided order ID.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        orderID path string true "Order ID"
// @Success      200  {object} domain.OrderResponse "Successfully retrieved order"
// @Failure      404  {object} domain.ErrorResponse  "Order not found"
// @Failure      500  {object} domain.ErrorResponse  "Internal server error"
// @Router       /order/{orderID} [get]
func (s *OrderService) GetOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("orderID")

	//Get Product API
	order, err := s.api.GetOrderByID(orderID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(
		domain.OrderResponse{
			StatusCode: 200,
			Message:    "Successfully retrieved order",
			Data:       order})
}

// DeleteOrder godoc
// @Summary      Delete order by ID
// @Description  Deletes an order based on the provided order ID.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        orderID path string true "Order ID"
// @Success      204  {object} domain.OrderResponse "Successfully deleted order"
// @Failure      404  {object} domain.ErrorResponse   "Order not found"
// @Failure      500  {object} domain.ErrorResponse   "Internal server error"
// @Router       /order/{orderID} [delete]
func (s *OrderService) DeleteOrder(c *fiber.Ctx) error {
	orderID := c.Params("orderID")

	//Delete Product API
	err := s.api.DeleteOrder(orderID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(204).JSON(
		domain.OrderResponse{
			StatusCode: 204,
			Message:    "Successfully Deleted order",
		})

}

func addUUID(uuid string, orderItems []domain.OrderItemRequest) []domain.OrderItem {
	var orders []domain.OrderItem

	// Convert the array of RequestItem to your struct type
	for _, item := range orderItems {
		// Convert each RequestItem to Product
		order := domain.OrderItem{
			OrderID:   uuid,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		// Append the struct to the struct array
		orders = append(orders, order)
	}

	return orders
}
