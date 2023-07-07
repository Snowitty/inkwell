package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowitty/inkwell/models"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{
		DB: db,
	}
}

func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	category := new(models.Categories)
	if err := ctx.BodyParser(category); err != nil {
		return err
	}

	// 创建分类记录并保存到数据库中
	if err := c.DB.Create(category).Error; err != nil {
		return err
	}

	return ctx.JSON(category)
}

func (c *CategoryController) GetCategories(ctx *fiber.Ctx) error {
	categories := []*models.Categories{}

	// 从数据库中获取所有分类记录
	if err := c.DB.Find(&categories).Error; err != nil {
		return err
	}

	return ctx.JSON(categories)
}

func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	category := new(models.Categories)

	// 从数据库中根据ID获取分类记录
	if err := c.DB.First(category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Category not found",
			})
		}
		return err
	}

	return ctx.JSON(category)
}

func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	category := new(models.Categories)
	if err := ctx.BodyParser(category); err != nil {
		return err
	}

	// 更新分类记录到数据库中
	if err := c.DB.Model(&models.Categories{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return err
	}

	return ctx.JSON(category)
}

func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// 从数据库中删除指定ID的分类记录
	if err := c.DB.Delete(&models.Categories{}, id).Error; err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Category deleted",
	})
}
