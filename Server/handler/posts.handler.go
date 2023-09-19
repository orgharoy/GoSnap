package handler

import "github.com/gofiber/fiber/v2"

func CreatePost(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User has created", "data": nil})
}
