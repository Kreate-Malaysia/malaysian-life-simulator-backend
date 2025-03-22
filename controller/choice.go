package controller

import (
	"gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChoiceController struct {
    ChoiceService *services.ChoiceService
}

func NewChoiceController(choiceService *services.ChoiceService) *ChoiceController {
    return &ChoiceController{ChoiceService: choiceService}
}

// Get all choices for a specific scenario
func (c *ChoiceController) GetChoices(ctx *gin.Context) {
	scenarioID, err := strconv.Atoi(ctx.Param("scenario_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scenario ID"})
		return
	}

	events, err := c.ChoiceService.GetChoices(scenarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}