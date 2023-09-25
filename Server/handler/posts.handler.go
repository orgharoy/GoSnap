package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/orgharoy/GoSnap/database"
	helperfunctions "github.com/orgharoy/GoSnap/helperFunctions"
	"github.com/orgharoy/GoSnap/model"
)

type InputPostStruct struct {
	Title       string `gorm:"varchar(255); not null" json:"title"`
	Description string `gorm:"varchar(255); not null" json:"description"`
	Image       string `gorm:"varchar(255)" json:"image"`
	Address     string `gorm:"null" json:"address"`
}

func CreatePosts(c *fiber.Ctx) error {

	db := database.DB

	var newPost InputPostStruct

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
	db := database.DB

	var posts []model.Post

	db.Find(&posts)

	if len(posts) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No Posts", "data": nil})
	}
	// return users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Posts Found", "data": posts})
}

func GetPost(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var post model.Post

	db.Find(&post, "id = ?", id)

	if post.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post not found", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Post Found", "data": post})
}

func GetMyPosts(c *fiber.Ctx) error {
	db := database.DB

	tokenString := c.Get("Authorization") // Assuming the token is sent in the Authorization header

	userID, err := helperfunctions.ExtractUserIDFromToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Parsing Token UID", "data": nil})
	}

	var posts []model.Post

	db.Find(&posts, "user_id = ?", userID)

	if len(posts) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No Posts Found", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Post Found", "data": posts})
}

func UpdatePost(c *fiber.Ctx) error {
	db := database.DB

	tokenString := c.Get("Authorization") // Assuming the token is sent in the Authorization header

	userId, err := helperfunctions.ExtractUserIDFromToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Parsing Token UID", "data": nil})
	}

	id := c.Params("id")

	var post model.Post

	db.Find(&post, "id = ?", id)

	if post.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Post of This ID Not Found", "data": nil})
	}

	var updatedPost InputPostStruct

	err = c.BodyParser(&updatedPost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	if updatedPost.Image == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No Image Found", "data": nil})
	}

	if post.UserID != userId {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "You Are Not the Author of This Post", "data": nil})
	}

	post.Title = updatedPost.Title
	post.Description = updatedPost.Description
	post.Image = updatedPost.Image
	post.Address = updatedPost.Address
	post.UpdatedAt = time.Now()

	db.Save(&post)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Successfully Updated The Post", "data": post})
}
