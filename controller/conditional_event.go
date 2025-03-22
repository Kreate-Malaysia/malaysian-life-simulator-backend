package controller

import (
	"gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConditionalEventController struct {
    ConditionalEventService *services.ConditionalEventService
}

func NewConditionalEventController(conditionalEventService *services.ConditionalEventService) *ConditionalEventController {
    return &ConditionalEventController{conditionalEventService}
}

// Get all conditional events for a specific scenario
func (c *ConditionalEventController) GetConditionalEvents(ctx *gin.Context) {
	scenarioID, err := strconv.Atoi(ctx.Param("scenario_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scenario ID"})
		return
	}

	events, err := c.ConditionalEventService.GetConditionalEvent(scenarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}