package routes

import (
	"bwastartup/config"
	"bwastartup/handlers"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
)

func Router() {
	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	var gin = gin.Default()
	api := gin.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.GET("/users", userHandler.LoginUser)
	gin.Run()
}
