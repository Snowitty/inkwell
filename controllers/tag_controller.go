// controllers/tags_controller.go
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowitty/inkwell/models"
	"gorm.io/gorm"
)

type TagsController struct {
	DB *gorm.DB
}

func NewTagsController(db *gorm.DB) *TagsController {
	return &TagsController{
		DB: db,
	}
}

func (c *TagsController) CreateTag(ctx *fiber.Ctx) error {
	tag := new(models.Tags)
	if err := ctx.BodyParser(tag); err != nil {
		return err
	}

	// 创建标签记录并保存到数据库中
	if err := c.DB.Create(tag).Error; err != nil {
		return err
	}

	return ctx.JSON(tag)
}

func (c *TagsController) GetTags(ctx *fiber.Ctx) error {
	tags := []*models.Tags{}

	// 从数据库中获取所有标签记录
	if err := c.DB.Find(&tags).Error; err != nil {
		return err
	}

	return ctx.JSON(tags)
}

func (c *TagsController) GetTagByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	tag := new(models.Tags)

	// 从数据库中根据ID获取标签记录
	if err := c.DB.First(tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Tag not found",
			})
		}
		return err
	}

	return ctx.JSON(tag)
}

func (c *TagsController) UpdateTag(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	tag := new(models.Tags)
	if err := ctx.BodyParser(tag); err != nil {
		return err
	}

	// 更新标签记录到数据库中
	if err := c.DB.Model(&models.Tags{}).Where("id = ?", id).Updates(tag).Error; err != nil {
		return err
	}

	return ctx.JSON(tag)
}

func (c *TagsController) DeleteTag(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// 从数据库中删除指定ID的标签记录
	if err := c.DB.Delete(&models.Tags{}, id).Error; err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Tag deleted",
	})
}
