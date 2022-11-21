package controllers

import (
	"encoding/json"
	"go-admin/src/models"
	"go-admin/src/services"

	"github.com/gofiber/fiber/v2"
)

func Ambassadors(c *fiber.Ctx) error {

	resp, err := services.UserService.Get("users",c.Cookies("jwt",""))

	if err != nil {
		return err
	}

	var ambassadors []models.User

	var users []models.User

	json.NewDecoder(resp.Body).Decode(&users)

	for _, user := range users{
		if user.IsAmbassador {
			ambassadors = append(ambassadors, user)
		}
	}

	return c.JSON(ambassadors)
}