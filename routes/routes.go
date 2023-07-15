package routes

import (
	"github.com/Snowitty/inkwell/controllers"
	"github.com/Snowitty/inkwell/models"
	"github.com/Snowitty/inkwell/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func authenticateUser(ctx *fiber.Ctx) error {
	//从请求中获取JWT令牌
	tokenString := ctx.Get("Authorization")

	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	//解析JWT令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("inkwell"), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	//处理解析错误
	if !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	//上下文中保存用户信息
	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))
	ctx.Locals("userID", userID)
	//继续下一个中间件
	return ctx.Next()
}

func accessControl() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		userID, ok := ctx.Locals("userID").(uint)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		//访问控制逻辑

		articles := []models.Articles{}
		if err := utils.DB.Where("user_id = ?", userID).Find(&articles).Error; err != nil {
			return err
		}

		ctx.Locals("articles", articles)

		return ctx.Next()
	}

}
func Setup(app *fiber.App, db *gorm.DB) {
	userController := controllers.NewUserController(db)
	tagController := controllers.NewTagsController(db)
	categoryController := controllers.NewCategoryController(db)
	articleController := controllers.NewArticleController(db)

	//用户路由
	app.Post("/users", userController.CreateUser)
	app.Get("/users", userController.GetUsers)
	app.Get("/users/:id", userController.CheckCurrentUser, userController.GetUserArticles, userController.GetUserByID)
	app.Put("/users/:id", userController.CheckAdmin, userController.UpdateUser)
	app.Delete("/users/:id", userController.CheckAdmin, userController.DeleteUser)
	app.Post("/register", userController.RegisterUser)
	app.Post("/login", authenticateUser, accessControl(), userController.Login)

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
}
