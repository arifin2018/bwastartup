package handlers

import (
	"bwastartup/campaign"
	"bwastartup/helpers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helpers.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.ApiResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {

	var input = campaign.GetCampaignDetailInput{}

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaignData, err := h.service.GetCampaignById(input)
	if err != nil {
		fmt.Println("sw2")
		response := helpers.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.ApiResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignData))
	c.JSON(http.StatusOK, response)

}
