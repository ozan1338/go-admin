package middlewares

import (
	"encoding/json"
	"go-admin/src/models"
	"go-admin/src/services"

	"github.com/gofiber/fiber/v2"
)


func IsAuthenticated(c *fiber.Ctx) error {

	resp, err := services.UserService.Get("user/admin", c.Cookies("jwt",""))

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"unauthenticated",
		})
	}

	var user  models.User

	json.NewDecoder(resp.Body).Decode(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"unauthenticated",
		})
	}

	c.Context().SetUserValue("user", user)

	return c.Next()
}