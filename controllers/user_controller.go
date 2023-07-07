// controllers/user_controller.go
package controllers

import (
	"github.com/Snowitty/inkwell/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(models.Users)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// 创建用户记录并保存到数据库中
	if err := c.DB.Create(user).Error; err != nil {
		return err
	}

	return ctx.JSON(user)
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	users := []*models.Users{}

	// 从数据库中获取所有用户记录
	if err := c.DB.Find(&users).Error; err != nil {
		return err
	}

	return ctx.JSON(users)
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := new(models.Users)

	// 从数据库中根据ID获取用户记录
	if err := c.DB.First(user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return err
	}

	return ctx.JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := new(models.Users)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// 更新用户记录到数据库中
	if err := c.DB.Model(&models.Users{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}

	return ctx.JSON(user)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// 从数据库中删除指定ID的用户记录
	if err := c.DB.Delete(&models.Users{}, id).Error; err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "User deleted",
	})
}
