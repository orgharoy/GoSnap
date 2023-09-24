package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orgharoy/GoSnap/handler"
	"github.com/orgharoy/GoSnap/middleware"
)

func Routes(app *fiber.App) {

	app.Post("/users", handler.CreateUser)
	app.Get("/users", middleware.AuthMiddleware, handler.GetUsers)
	app.Get("/user/:id", middleware.AuthMiddleware, handler.GetUser)
	app.Put("/user/:id", middleware.AuthMiddleware, handler.UpdateUser)
	app.Delete("/user/:id", middleware.AuthMiddleware, handler.DeleteUser)

	app.Post("/login", handler.Login)

	app.Post("/posts", middleware.AuthMiddleware, handler.CreatePosts)
	app.Get("/posts", middleware.AuthMiddleware, handler.GetPosts)
}
