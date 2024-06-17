package routes

import (
	"bwastartup/auth"
	"bwastartup/auth/middlewares"
	"bwastartup/config"
	"bwastartup/handlers"
	"bwastartup/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router() {
	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	userHandler := handlers.NewUserHandler(userService, authService)

	tokenValidate, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.z0Sl9bmutbXwvQpTxua76AA5G8oGGqqqO0GKF4eOAJ4")
	if err != nil {
		panic(err.Error())
	}

	if tokenValidate.Valid {
		fmt.Println("valid")
	} else {
		fmt.Println("not valid")
	}

	var gin = gin.Default()
	api := gin.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.LoginUser)
	api.POST("/avatars", middlewares.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	gin.Run()
}
