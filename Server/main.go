package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/orgharoy/GoSnap/database"
	"github.com/orgharoy/GoSnap/routes"
)

func main() {

	err := database.Connect()

	if err != nil {
		fmt.Println(err, " failed to connect to database")
	}

	app := fiber.New()

	app.Use(cors.New())

	routes.Routes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")
}
