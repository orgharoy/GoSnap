package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/orgharoy/GoSnap/database"
	"github.com/orgharoy/GoSnap/model"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `gorm:"varchar(255); not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

func Login(c *fiber.Ctx) error {
	db := database.DB

	user := new(LoginInput)

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	var existingUser model.User

	db.Find(&existingUser, "email = ?", user.Email)

	if existingUser.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No User Found", "data": nil})
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Passwords do not match", "data": nil})
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = existingUser.ID
	claims["exp"] = now.Add(1200).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte("poiwnskjfowijef8972POIUYhqawiehjis"))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}
