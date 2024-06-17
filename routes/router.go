package routes

import (
	"bwastartup/auth"
	"bwastartup/auth/middlewares"
	"bwastartup/campaign"
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

	campaignRepository := campaign.NewRepository(config.DB)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handlers.NewCampaignHandler(campaignService)

	var gin = gin.Default()
	api := gin.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.LoginUser)
	api.POST("/avatars", middlewares.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaign", campaignHandler.GetCampaigns)

	gin.Run()
}
