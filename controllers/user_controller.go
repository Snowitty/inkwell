package controllers

import (
	"github.com/Snowitty/inkwell/models"
	"github.com/Snowitty/inkwell/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	if err := utils.DB.Create(&user).Error; err != nil {
		return err
	}

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := new(models.User)
	if err := utils.DB.First(&user, id).Error; err != nil {
		return err
	}

	return c.JSON(user)
}
