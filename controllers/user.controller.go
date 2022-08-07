package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/Muhammad5943/go-fiber-gorm/database"
	"github.com/Muhammad5943/go-fiber-gorm/models/entity"
	"github.com/Muhammad5943/go-fiber-gorm/models/request"
	"github.com/Muhammad5943/go-fiber-gorm/models/response"
	"github.com/Muhammad5943/go-fiber-gorm/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []entity.User

	// err := database.DB.Debug().Find(&users) // Debug() used to show the query
	err := database.DB.Find(&users)
	if err.Error != nil {
		fmt.Println(err.Error)
	}

	if len(users) == 0 {
		fmt.Println("No user exist")
		return c.Status(404).SendString("User Not Found")
	}

	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	// initial request for create user and filter using body parser
	user := new(request.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	// validate User using third party
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "you might have bad request",
			"error":   strings.Split(errValidate.Error(), "Key:"),
		})
	}

	// Create User
	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
	}

	// hashing password
	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Fatal("Error :", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	// input heshed password to data
	newUser.Password = hashedPassword

	// create user
	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		fmt.Println("Failed store user")
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed store user",
		})
	}

	// return response
	return c.Status(201).JSON(fiber.Map{
		"message": "success to store new user",
		"user":    newUser,
	})
}

func GetUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User by userId " + userId + " not found, error: " + err.Error(),
		})
	}

	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success To get User " + userResponse.Name,
		"user":    userResponse,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user has bad request",
		})
	}

	// user data from database or table
	var user entity.User

	userId := c.Params("userId")

	//Check Available User
	err := database.DB.First(&user, "id = ?", userId).Error
	fmt.Println("err ", err)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// UPDATE USER DATA
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// Update User Data
	return c.Status(202).JSON(fiber.Map{
		"message": "Success Updated the Request",
		"data":    user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	var user entity.User

	// Check available user
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "user not found",
			"error":   err,
		})
	}

	errDeleted := database.DB.Debug().Delete(&user).Error
	if errDeleted != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "User Was Deleted",
	})
}
