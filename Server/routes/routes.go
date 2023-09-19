package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orgharoy/GoSnap/handler"
)

func Routes(app *fiber.App) {

	app.Post("/users", handler.CreateUser)
	app.Get("/users", handler.GetUsers)
	app.Get("/user/:id", handler.GetUser)
	app.Put("/user/:id", handler.UpdateUser)
	app.Delete("/user/:id", handler.DeleteUser)

	app.Post("/login", handler.Login)

	app.Post("/posts", handler.CreatePost)
}
