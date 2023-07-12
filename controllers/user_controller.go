// controllers/user_controller.go
package controllers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Snowitty/inkwell/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

func (c *UserController) RegisterUser(ctx *fiber.Ctx) error {
	user := new(models.Users)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// 检查用户名是否已存在
	existingUser := new(models.Users)
	if err := c.DB.Where("username = ?", user.Username).First(existingUser).Error; err == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	// 对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// 创建用户记录并保存到数据库中
	if err := c.DB.Create(user).Error; err != nil {
		return err
	}

	return ctx.JSON(user)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	loginData := new(models.LoginData)
	if err := ctx.BodyParser(loginData); err != nil {
		return err
	}

	// 查找用户
	user := new(models.Users)
	if err := c.DB.Where("username = ?", loginData.Username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid credentials",
			})
		}
		return err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// 生成 JWT
	token, err := generateJWT(user.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func generateJWT(userID uint) (string, error) {
	// 创建声明
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 过期时间为24小时
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名字符串
	tokenString, err := token.SignedString([]byte("inkwell")) // 替换为你的密钥
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenString, nil
}

func (c *UserController) GetCurrentUser(ctx *fiber.Ctx) (*models.Users, error) {

	userID, ok := ctx.Locals("userID").(uint)
	if !ok {
		return nil, fmt.Errorf("user ID not found in context")

	}

	user := new(models.Users)
	if err := c.DB.First(user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}

		return nil, err
	}

	return user, nil

}

func (c *UserController) CheckAdmin(ctx *fiber.Ctx) error {
	user, err := c.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get current user",
		})
	}

	if !user.IsAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access denied",
		})
	}

	return ctx.Next()
}

func (c *UserController) CheckCurrentUser(ctx *fiber.Ctx) error {

	//获取当前登录用户
	user, err := c.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get current user",
		})
	}
	//从路由参数中获取用户ID
	userIDParam := ctx.Params("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}
	//检查当前用户
	if user.ID != uint(userID) && !user.IsAdmin {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access denied",
		})
	}

	return ctx.Next()

}
