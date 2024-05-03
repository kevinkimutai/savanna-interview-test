package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) OrderRouter(api fiber.Router) {
	api.Post("/", s.auth.IsAuthenticated, s.order.CreateOrder)
	api.Get("/", s.auth.IsAuthenticated, s.auth.AllowedRoles("Admin"), s.order.GetAllOrders)
	api.Get("/:orderID", s.auth.IsAuthenticated, s.order.GetOrderByID)
	api.Delete("/:orderID", s.auth.IsAuthenticated, s.auth.AllowedRoles("Admin"), s.order.DeleteOrder)

}
