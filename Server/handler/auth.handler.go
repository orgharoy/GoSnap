package handler

import (
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

	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()

	// Set claims (payload data) in the token.
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = existingUser.ID
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour * 24).Unix()

	// Sign the token with the secret key.
	tokenString, err := token.SignedString([]byte("pwoEQuF2jdk4c!nW$Nuew^rf6kjnV"))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": "generating JWT Token failed"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}
