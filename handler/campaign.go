package handler

import (
	"funding/campaign"
	"funding/helper"
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
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)

	if len(campaigns) == 0 {
		response := helper.APIResponse("Campaigns not found", http.StatusNotFound, "Success", formatter)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("List Campaign", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}
