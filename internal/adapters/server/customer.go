package server

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func (s *ServerAdapter) CustomerRouter(api fiber.Router) {
// 	api.Post("/", s.auth.IsAuthenticated, s.customer.CreateCustomer)
// 	api.Get("/", s.auth.IsAuthenticated, s.auth.AllowedRoles("Admin"), s.customer.GetAllCustomers)
// 	api.Get("/:customerID", s.auth.IsAuthenticated, s.auth.AllowedRoles("Admin"), s.customer.GetCustomerByID)
// 	api.Delete("/:customerID", s.auth.IsAuthenticated, s.auth.AllowedRoles("Admin"), s.customer.DeleteCustomer)

// }
