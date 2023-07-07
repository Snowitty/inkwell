// controllers/article_controller.go
package controllers

import (
	"github.com/Snowitty/inkwell/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ArticleController struct {
	DB *gorm.DB
}

func NewArticleController(db *gorm.DB) *ArticleController {
	return &ArticleController{
		DB: db,
	}
}

func (c *ArticleController) CreateArticle(ctx *fiber.Ctx) error {
	article := new(models.Articles)
	if err := ctx.BodyParser(article); err != nil {
		return err
	}

	// 创建文章记录并保存到数据库中
	if err := c.DB.Create(article).Error; err != nil {
		return err
	}

	return ctx.JSON(article)
}

func (c *ArticleController) GetArticles(ctx *fiber.Ctx) error {
	articles := []*models.Articles{}

	// 从数据库中获取所有文章记录
	if err := c.DB.Find(&articles).Error; err != nil {
		return err
	}

	return ctx.JSON(articles)
}

func (c *ArticleController) GetArticleByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	article := new(models.Articles)

	// 从数据库中根据ID获取文章记录
	if err := c.DB.First(article, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Article not found",
			})
		}
		return err
	}

	return ctx.JSON(article)
}

func (c *ArticleController) UpdateArticle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	article := new(models.Articles)
	if err := ctx.BodyParser(article); err != nil {
		return err
	}

	// 更新文章记录到数据库中
	if err := c.DB.Model(&models.Articles{}).Where("id = ?", id).Updates(article).Error; err != nil {
		return err
	}

	return ctx.JSON(article)
}

func (c *ArticleController) DeleteArticle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// 从数据库中删除指定ID的文章记录
	if err := c.DB.Delete(&models.Articles{}, id).Error; err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Article deleted",
	})
}
