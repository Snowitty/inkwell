package routes

import (
	"github.com/Snowitty/inkwell/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	userController := controllers.NewUserController(db)
	tagController := controllers.NewTagsController(db)
	categoryController := controllers.NewCategoryController(db)
	articleController := controllers.NewArticleController(db)

	//用户路由
	app.Post("/users", userController.CreateUser)
	app.Get("/users", userController.GetUsers)
	app.Get("/users/:id", userController.GetUserByID)
	app.Put("/users/:id", userController.UpdateUser)
	app.Delete("/users/:id", userController.DeleteUser)
	app.Post("/register", userController.RegisterUser)
	app.Post("/login", userController.Login)

	// 标签路由
	app.Post("/tags", tagController.CreateTag)
	app.Get("/tags", tagController.GetTags)
	app.Get("/tags/:id", tagController.GetTagByID)
	app.Put("/tags/:id", tagController.UpdateTag)
	app.Delete("/tags/:id", tagController.DeleteTag)

	// 分类路由
	app.Post("/categories", categoryController.CreateCategory)
	app.Get("/categories", categoryController.GetCategories)
	app.Get("/categories/:id", categoryController.GetCategoryByID)
	app.Put("/categories/:id", categoryController.UpdateCategory)
	app.Delete("/categories/:id", categoryController.DeleteCategory)

	// 文章路由
	app.Post("/articles", articleController.CreateArticle)
	app.Get("/articles", articleController.GetArticles)
	app.Get("/articles/:id", articleController.GetArticleByID)
	app.Put("/articles/:id", articleController.UpdateArticle)
	app.Delete("/articles/:id", articleController.DeleteArticle)

	// 应用中间件
	app.Use(userController.AuthenticateUser)
	app.Use(userController.AccessControl())

}
