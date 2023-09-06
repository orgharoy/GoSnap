package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orgharoy/GoSnap/handler"
)

func Routes(app *fiber.App) {
	app.Get("/", handler.HelloWorld)

	app.Post("/users", handler.CreateUser)
	app.Get("/users", handler.GetUsers)
	app.Get("/user/:id", handler.GetUser)

	app.Post("/login", handler.Login)
}
