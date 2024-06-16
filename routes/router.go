package routes

import (
	"bwastartup/auth"
	"bwastartup/config"
	"bwastartup/handlers"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
)

func Router() {
	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	userHandler := handlers.NewUserHandler(userService, authService)

	var gin = gin.Default()
	api := gin.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.GET("/email_checkers", userHandler.LoginUser)
	api.POST("/avatars", userHandler.UploadAvatar)
	gin.Run()
}
