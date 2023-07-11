package main

import (
	"log"

	"github.com/Snowitty/inkwell/routes"
	"github.com/Snowitty/inkwell/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {

	err := utils.InitDB("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// 路由设置
	routes.Setup(app, utils.DB)

	// 注册 Swagger 中间件
	app.Get("/swagger/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./docs/swagger.json")
	})

	app.Static("/swagger", "./docs")

	app.Listen(":3000")
}
