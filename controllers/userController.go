package controllers

import (
	"bwastartup/config"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	// var user = user.User{
	// 	Name:           "azriel",
	// 	Occupation:     "Programmer",
	// 	Email:          "azrielgdr@gmail.com",
	// 	Password:       "password",
	// 	AvatarFileName: "azriel.jpg",
	// 	Token:          "awd",
	// 	CreatedAt:      time.Now(),
	// 	UpdatedAt:      time.Now(),
	// }
	// userRepository.Save(user)
	// c.JSON(http.StatusAccepted, user)
	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	userInput := user.RegisterUser{
		Name:       "rafiq",
		Occupation: "siswa",
		Email:      "rafiq@lenna.ai",
		Password:   "rafiq",
		Role:       "user",
	}
	registerUserInput, err := userService.RegisterUser(userInput)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusAccepted, registerUserInput)

}
