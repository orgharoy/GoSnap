package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/orgharoy/GoSnap/database"
	helperfunctions "github.com/orgharoy/GoSnap/helperFunctions"
	"github.com/orgharoy/GoSnap/model"
)

type CreatePostStruct struct {
	Title       string `gorm:"varchar(255); not null" json:"title"`
	Description string `gorm:"varchar(255); not null" json:"description"`
	Image       string `gorm:"varchar(255)" json:"image"`
	Address     string `gorm:"null" json:"address"`
}

func CreatePosts(c *fiber.Ctx) error {

	db := database.DB

	var newPost CreatePostStruct

	tokenString := c.Get("Authorization") // Assuming the token is sent in the Authorization header

	userID, err := helperfunctions.ExtractUserIDFromToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Parsing Token UID", "data": nil})
	}

	err = c.BodyParser(&newPost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	if newPost.Image == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "No Image Found", "data": nil})
	}

	var post model.Post

	post.Title = newPost.Title
	post.Description = newPost.Description
	post.Image = newPost.Image
	post.Address = newPost.Address
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	post.UserID = userID

	err = db.Create(&post).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not create Post", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Post has been created", "data": post})
}

func GetPosts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Post has been created", "data": "post"})
}
