package controllers

import (
	"github.com/Snowitty/inkwell/models"
	"github.com/Snowitty/inkwell/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)

	if err := c.BodyParser(tag); err != nil {
		return err
	}

	if err := utils.DB.Create(&tag).Error; err != nil {
		return err
	}

	return c.JSON(tag)
}

func GetTag(c *fiber.Ctx) error {
	id := c.Params("id")

	tag := new(models.Tag)
	if err := utils.DB.First(&tag, id).Error; err != nil {
		return err
	}

	return c.JSON(tag)
}

func DeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")

	tag := new(models.Tag)
	if err := utils.DB.First(&tag, id).Error; err != nil {
		return err
	}

	if err := utils.DB.Delete(&tag).Error; err != nil {
		return err
	}

	return c.SendString("tag删除成功")
}
