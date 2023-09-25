package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orgharoy/GoSnap/handler"
	"github.com/orgharoy/GoSnap/middleware"
)

func Routes(app *fiber.App) {

	app.Post("/user", handler.CreateUser)                                  //-> Create User
	app.Get("/users", middleware.AuthMiddleware, handler.GetUsers)         //-> Get All Users
	app.Get("/user/:id", middleware.AuthMiddleware, handler.GetUser)       //-> Get User By ID
	app.Put("/user/:id", middleware.AuthMiddleware, handler.UpdateUser)    //-> Update User By ID
	app.Delete("/user/:id", middleware.AuthMiddleware, handler.DeleteUser) //-> Delete User By ID

	app.Post("/login", handler.Login) //-> Login Handler

	app.Post("/posts", middleware.AuthMiddleware, handler.CreatePosts)  //-> Create Post
	app.Get("/posts", middleware.AuthMiddleware, handler.GetPosts)      //-> Get All Posts
	app.Get("/post/:id", middleware.AuthMiddleware, handler.GetPost)    //-> Get Post By ID
	app.Get("/my-posts", middleware.AuthMiddleware, handler.GetMyPosts) //-> Get All Post By A User (Getting User ID From JWT)
	app.Put("/post/:id", middleware.AuthMiddleware, handler.UpdatePost) //-> Update Post By ID

}
