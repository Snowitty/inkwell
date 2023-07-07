package main

import (
	"log"

	"github.com/Snowitty/inkwell/routes"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局身份验证中间件
func authenticateUser(ctx *fiber.Ctx) error {
	// 从请求中获取 JWT 令牌
	tokenString := ctx.Get("Authorization")

	// 解析 JWT 令牌
	token, err := parseJWTToken(tokenString)
	if err != nil || !token.Valid {
		// 身份验证失败，返回未授权错误
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// 在上下文中存储用户身份信息，以便后续使用
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["userID"].(string)
	ctx.Locals("userID", userID)

	// 继续执行下一个中间件或路由处理函数
	return ctx.Next()
}

// 解析 JWT 令牌
func parseJWTToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证密钥并返回
		return []byte("x6HSPyVdVb31WodHyIIayiQM7S3NrH6lZ1UxwVgEFNk="), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// 全局访问控制中间件
func accessControl() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// 访问控制逻辑...
		return ctx.Next()
	}
}

func main() {
	app := fiber.New()

	// 读取数据库配置
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// 获取数据库连接配置
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.username")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")

	// 构建数据库连接字符串
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// 注册身份验证中间件
	app.Use(authenticateUser)

	// 注册访问控制中间件
	app.Use(accessControl())

	// 路由设置
	routes.Setup(app, db)

	// 注册 Swagger 中间件
	app.Get("/swagger/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./docs/swagger.json")
	})

	app.Static("/swagger", "./docs")

	app.Listen(":3000")
}
