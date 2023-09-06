package handler

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/orgharoy/GoSnap/database"
	helperfunctions "github.com/orgharoy/GoSnap/helperFunctions"
	"github.com/orgharoy/GoSnap/model"
	"golang.org/x/crypto/bcrypt"
)

func HelloWorld(c *fiber.Ctx) error {
	//db :=
	return c.Status(200).JSON("Hello")
	//c.Send("Hello, World!")
}

func CreateUser(c *fiber.Ctx) error {

	db := database.DB

	user := new(model.User)

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	//data validation
	// -> empty fields
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Empty Fields", "data": nil})
	}

	// -> email validation
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegex.MatchString(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid Email", "data": nil})
	}

	if !helperfunctions.IsValidPassword(user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Password must be longer than 8 characters, contain both uppercase and lowercase letter and have have special characters", "data": nil})
	}

	//to check if user exists
	var existingUser model.User

	db.Find(&existingUser, "email = ?", user.Email)

	if existingUser.ID != uuid.Nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "User Already Exists", "data": nil})
	}
	
	//hashing password

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error Hashing Password", "data": err})
	}

	user.Password = string(bytes)

	//creating a new user
	err = db.Create(&user).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	// Return the created user
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

func GetUsers(c *fiber.Ctx) error {

	db := database.DB

	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	// return users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

func GetUser(c *fiber.Ctx) error {

	db := database.DB

	id := c.Params("id")

	var user model.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}
