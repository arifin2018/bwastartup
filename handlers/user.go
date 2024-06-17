package handlers

import (
	"bwastartup/auth"
	"bwastartup/helpers"
	"bwastartup/user"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input = new(user.RegisterUser)
	err := c.ShouldBindBodyWithJSON(input)
	if err != nil {
		var errorsData []string
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, v := range err.(validator.ValidationErrors) {
				errorsData = append(errorsData, v.Error())
			}
		}

		errorMessage := gin.H{
			"errors": errorsData,
		}

		response := helpers.ApiResponse("Account failed", http.StatusOK, "success", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userRegister, err := h.userService.RegisterUser(*input)
	if err != nil {
		var errorsData []string
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// It's a validation error, iterate through each error
			for _, e := range validationErrs {
				// Cast the FieldError into our ValidationError and append to the slice
				errorsData = append(errorsData, e.Error())
			}
		} else {
			// Handle other types of errors here, for example, *mysql.MySQLError
			errorsData = append(errorsData, err.Error())
		}
		errorMessage := gin.H{
			"errors": errorsData,
		}

		response := helpers.ApiResponse("Account failed", http.StatusOK, "success", errorMessage)
		fmt.Println(response)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// userRegister.Id
	token, err := h.authService.GenerateKey(int(userRegister.Id))
	if err != nil {
		errorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helpers.ApiResponse("Account failed", http.StatusOK, "success", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatUser := user.FormatUser(userRegister, token)

	response := helpers.ApiResponse("Account has been registered", http.StatusOK, "success", formatUser)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input = new(user.LoginUser)
	if err := c.ShouldBindBodyWithJSON(input); err != nil {
		var errorsData []string
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, v := range err.(validator.ValidationErrors) {
				errorsData = append(errorsData, v.Error())
			}
		}

		errorMessage := gin.H{
			"errors": errorsData,
		}

		response := helpers.ApiResponse("Account failed", http.StatusOK, "success", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.LoginUser(*input)
	if err != nil {
		errorMessage := gin.H{
			"errors": err.Error(),
		}

		response := helpers.ApiResponse("Login failed", http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateKey(int(loggedinUser.Id))
	if err != nil {
		response := helpers.ApiResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helpers.ApiResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helpers.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helpers.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helpers.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		fmt.Println(err.Error())
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.Id

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponse("Failed to upload avatar image2", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(int(userID), path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponse("Failed to upload avatar image3", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helpers.ApiResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
