package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type CreatePostStruct struct {
	Title       string `gorm:"varchar(255); not null" json:"title"`
	Description string `gorm:"varchar(255); not null" json:"description"`
	Image       string `gorm:"varchar(255)" json:"image"`
	Address     string `gorm:"null" json:"address"`
}

func CreatePost(c *fiber.Ctx) error {

	//db := database.DB

	//var newPost CreatePostStruct

	tokenString := c.Get("Authorization") // Assuming the token is sent in the Authorization header

	fmt.Println(tokenString)

	// Parse the token using the secret key.
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
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Validatin error", "data": nil})

			}
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Default", "data": err})
	}

	// Check if the token is valid.
	if !token.Valid {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Invalid Token", "data": nil})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the user ID from the claims.
		userId, ok := claims["userID"]
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Parsing Token UID", "data": userId})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": userId})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "last step", "data": nil})

	// Parse the JWT token
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Check the signing method, which should match the one used during token creation
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, jwt.ErrSignatureInvalid
	// 	}
	// 	// Return the secret key used for signing
	// 	return []byte("pwoEQuF2jdk4c!nW$Nuew^rf6kjnV"), nil
	// })

	// if err != nil || !token.Valid {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	// }

	// // Token is valid, you can now access claims
	// claims, _ := token.Claims.(jwt.MapClaims)

	// // Extract the user ID from claims
	// userIDStr := claims["userID"].(string)

	// userID, err := uuid.Parse(userIDStr)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "cannot covert to uuid", "data": nil})
	// }

	/////////////////////////////////////////////////////////////////////////

	// err = c.BodyParser(newPost)

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	// }

	// if newPost.Image == "" {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "No Image Found", "data": nil})
	// }

	// var post model.Post

	// post.Title = newPost.Title
	// post.Description = newPost.Description
	// post.Image = newPost.Image
	// post.Address = newPost.Address
	// post.CreatedAt = time.Now()
	// post.UpdatedAt = time.Now()
	// post.UserID = userID

	// err = db.Create(&post).Error

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not create Post", "data": err})
	// }

	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Post has been created", "data": post})
}
