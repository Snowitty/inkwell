package main

import (
	"log"

	"github.com/Snowitty/inkwell/models"
	"github.com/Snowitty/inkwell/routes"
	"github.com/Snowitty/inkwell/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// 全局身份验证中间件
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

// 全局访问控制中间件
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
func main() {

	err := utils.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(authenticateUser)
	app.Use(accessControl())

	// 路由设置
	routes.Setup(app, utils.DB)

	// 注册 Swagger 中间件
	app.Get("/swagger/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./docs/swagger.json")
	})

	app.Static("/swagger", "./docs")

	app.Listen(":3000")
}
