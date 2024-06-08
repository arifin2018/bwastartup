package handlers

import (
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
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{
		userService: userService,
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

	formatUser := user.FormatUser(userRegister, "token")

	response := helpers.ApiResponse("Account has been registered", http.StatusOK, "success", formatUser)

	c.JSON(http.StatusOK, response)
}
