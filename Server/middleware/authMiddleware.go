package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/orgharoy/GoSnap/database"
	"github.com/orgharoy/GoSnap/model"
)

func AuthMiddleware(c *fiber.Ctx) error {

	db := database.DB

	var user model.User

	tokenString := c.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("pwoEQuF2jdk4c!nW$Nuew^rf6kjnV"), nil
	})

	if err != nil {
		// Handle the error and provide additional information.
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			if validationErr.Errors&jwt.ValidationErrorMalformed != 0 {
				// Token is not valid JWT format.
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Malformed Token", "data": nil})

			} else if validationErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is expired or not yet valid.
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Token Expired", "data": nil})

			} else {
				// Other validation errors.
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Validation error", "data": nil})

			}
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Validating Token", "data": err})
	}

	if !token.Valid {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Invalid Token", "data": nil})
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Parsing Claims", "data": nil})
	}

	userId, ok := claims["userID"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Parsing Token UID", "data": nil})
	}

	userID, err := uuid.Parse(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Parsing ID String to UUID", "data": nil})
	}

	db.Find(&user, "id = ?", userID)

	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No User Found", "data": nil})
	}

	return c.Next()
}
