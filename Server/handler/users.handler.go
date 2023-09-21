package handler

import (
	"errors"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/orgharoy/GoSnap/database"
	helperfunctions "github.com/orgharoy/GoSnap/helperFunctions"
	"github.com/orgharoy/GoSnap/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserResponse struct {
	ID             uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primary_key" json:"id"`
	FirstName      string    `gorm:"varchar(255); not null" json:"firstName"`
	LastName       string    `gorm:"varchar(255); not null" json:"lastName"`
	Email          string    `gorm:"varchar(255); not null" json:"email"`
	ProfilePicture string    `gorm:"varchar(255)" json:"profilePicture"`
	Bio            string    `gorm:"null" json:"bio"`
	Address        string    `gorm:"null" json:"address"`
	CreatedAt      time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"not null" json:"updatedAt"`
}

type UpdateUserValues struct {
	FirstName      string `gorm:"varchar(255); not null" json:"firstName"`
	LastName       string `gorm:"varchar(255); not null" json:"lastName"`
	Email          string `gorm:"varchar(255); not null" json:"email"`
	ProfilePicture string `gorm:"varchar(255)" json:"profilePicture"`
	Bio            string `gorm:"null" json:"bio"`
	Address        string `gorm:"null" json:"address"`
}

func CreateResponseUser(userModel model.User) UserResponse {
	return UserResponse{
		ID:             userModel.ID,
		FirstName:      userModel.FirstName,
		LastName:       userModel.LastName,
		Email:          userModel.Email,
		ProfilePicture: userModel.ProfilePicture,
		Bio:            userModel.Bio,
		Address:        userModel.Address,
		CreatedAt:      userModel.CreatedAt,
		UpdatedAt:      userModel.UpdatedAt,
	}
}

func CreateUser(c *fiber.Ctx) error {

	db := database.DB

	var user model.User

	err := c.BodyParser(&user)

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

	// -> password validation
	if !helperfunctions.IsValidPassword(user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Password must be longer than 8 characters, contain both uppercase and lowercase letter and have have special characters", "data": nil})
	}

	//to check if user exists
	var existingUser model.User

	db.Find(&existingUser, "email = ?", user.Email)

	if existingUser.ID != uuid.Nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "error", "message": "Account with this Email already exists", "data": nil})
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

	responseUser := CreateResponseUser(user)

	// Return the created user
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User has created", "data": responseUser})
}

func GetUsers(c *fiber.Ctx) error {

	db := database.DB

	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}

	responseUsers := []UserResponse{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	// return users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": responseUsers})
}

func GetUser(c *fiber.Ctx) error {

	db := database.DB

	id := c.Params("id")

	var user model.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	responseUser := CreateResponseUser(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Found", "data": responseUser})
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var user model.User

	var updateUser UpdateUserValues

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": updateUser})
	}

	password := c.Get("password")

	//data validation

	// -> empty fields
	if updateUser.FirstName == "" || updateUser.LastName == "" || updateUser.Email == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Empty Fields", "data": nil})
	}

	// -> email validation
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegex.MatchString(updateUser.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid Email", "data": nil})
	}

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	if updateUser.Email != user.Email {
		var existingUser model.User
		db.Where("email = ? AND id != ?", updateUser.Email, id).Find(&existingUser)

		if existingUser.ID != uuid.Nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "error", "message": "Email Already Exists", "data": nil})
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Passwords do not match", "data": nil})
	}

	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName
	user.Email = updateUser.Email
	user.ProfilePicture = updateUser.ProfilePicture
	user.Bio = updateUser.Bio
	user.Address = updateUser.Address
	user.UpdatedAt = time.Now()

	db.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Successfully Updated The User", "data": user})

}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")
	password := c.Get("password")

	var user model.User

	// Find the user by ID
	if err := db.Find(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
		}
		// Handle other database errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Database error", "data": nil})
	}

	// Check if the provided password matches the user's password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Password incorrect", "data": nil})
	}

	// Delete the user
	if err := db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User deleted successfully", "data": nil})
}
