package controller

import (
	"gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RandomEventController struct {
    RandomEventService *services.RandomEventService
}

func NewRandomEventController(randomEventService *services.RandomEventService) *RandomEventController {
    return &RandomEventController{randomEventService}
}

// Get all random events for a specific scenario
func (r *RandomEventController) GetRandomEvents(ctx *gin.Context) {
	scenarioID, err := strconv.Atoi(ctx.Param("scenario_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scenario ID"})
		return
	}

	events, err := r.RandomEventService.GetRandomEvents(scenarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (r *RandomEventController) RollRandomEvent(ctx *gin.Context) {
	scenarioID, err := strconv.Atoi(ctx.Param("scenario_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scenario ID"})
		return
	}

	events, err := r.RandomEventService.GetRandomEvents(scenarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	event := r.RandomEventService.SelectRandomEvent(events)

	ctx.JSON(http.StatusOK, event)
}