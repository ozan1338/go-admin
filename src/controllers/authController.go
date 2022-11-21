package controllers

import (
	"encoding/json"
	"go-admin/src/models"
	"go-admin/src/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	data["is_ambassador"] = strconv.FormatBool(false)
	
	resp, err := services.UserService.Post("register", "", data)

	if err != nil {
		return err
	}

	var user models.User

	json.NewDecoder(resp.Body).Decode(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	data["scope"] = "admin"
	
	resp, err := services.UserService.Post("login", "", data)
	if err != nil {
		return err
	}

	var respone map[string]string

	json.NewDecoder(resp.Body).Decode(&respone)


	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    respone["jwt"],
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {

	return c.JSON(c.Context().UserValue("user"))
}

func Logout(c *fiber.Ctx) error {
	services.UserService.Post("logout", c.Cookies("jwt",""), nil)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	resp, err := services.UserService.Put("users/info", c.Cookies("jwt",""), data)

	if err != nil {
		return err
	}

	var user models.User

	json.NewDecoder(resp.Body).Decode(&user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	resp, err := services.UserService.Put("users/password", c.Cookies("jwt",""), data)

	if err != nil {
		return err
	}

	var user models.User

	json.NewDecoder(resp.Body).Decode(&user)

	return c.JSON(user)
}
