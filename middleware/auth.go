// package middleware

// import (
// 	"strings"
// 	"github.com/gofiber/fiber/v2"
// )

// func NewMiddleware() fiber.Handler {
// 	return AuthMiddleware
// }

// func AuthMiddleware(c *fiber.Ctx) error {
// 	session, err := store.Get(c)

// 	if strings.Split(c.Path(), "auth")[0] == "login" {
// 		return c.Next()
// 	}

// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "not authorized",
// 		})
// 	}

// 	if session.Get(AUTH_key) == nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "not authorized",
// 		})
// 	}
// 	return c.Next()
// }