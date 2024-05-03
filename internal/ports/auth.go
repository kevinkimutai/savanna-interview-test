package ports

import "github.com/gofiber/fiber/v2"

type AuthHandlerPort interface {
	IsAuthenticated(c *fiber.Ctx) error
	AllowedRoles(admin string) func(c *fiber.Ctx) error
}
