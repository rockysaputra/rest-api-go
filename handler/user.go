package handler

import (
	"belajar-gorm/helper"
	"belajar-gorm/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	input := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser)
	response := helper.APIResponse("Account Has been created", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	input := user.LoginUserInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	findUser, err := h.userService.LoginUser(input)
	if err != nil {
		response := helper.APIResponse("Username or Password is wrong", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(findUser)
	response := helper.APIResponse("Login Success", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}
