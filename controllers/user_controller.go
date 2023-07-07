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

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := new(models.User)
	if err := utils.DB.First(&user, id).Error; err != nil {
		return err
	}

	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return err
	}

	user.Username = newUser.Username
	user.Password = newUser.Password
	user.Email = newUser.Email
	user.NickName = newUser.NickName
	user.Avatar = newUser.Avatar

	if err := utils.DB.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(user)

}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	if err := utils.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := utils.DB.Delete(&user).Error; err != nil {
		return err
	}

	return c.SendString("user删除成功")
}
