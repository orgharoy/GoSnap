package helperfunctions

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func IsValidPassword(password string) bool {
	// Check for at least 8 characters
	if len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter
	uppercaseRegex := regexp.MustCompile("[A-Z]")
	if !uppercaseRegex.MatchString(password) {
		return false
	}

	// Check for at least one lowercase letter
	lowercaseRegex := regexp.MustCompile("[a-z]")
	if !lowercaseRegex.MatchString(password) {
		return false
	}

	// Check for at least one special character (you can modify the character class)
	specialCharRegex := regexp.MustCompile("[!@#$%^&*()_+]")
	return specialCharRegex.MatchString(password)
}

// ExtractUserIDFromToken extracts the user ID from a JWT token string.

func ExtractUserIDFromToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("pwoEQuF2jdk4c!nW$Nuew^rf6kjnV"), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return uuid.Nil, err
	}

	userIDStr, ok := claims["userID"].(string)
	if !ok {
		return uuid.Nil, fiber.ErrForbidden // userID not found in claims
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, err // Failed to parse userID as UUID
	}

	return userID, nil
}
