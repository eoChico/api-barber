package handler

import (
	"net/http"

	"github.com/eoChico/api-barber/models"
	"github.com/gin-gonic/gin"
)

func CreateScheduling(ctx *gin.Context) {
	request := CreateSchedulingRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	scheduling := models.Scheduling{
		Description: request.Description,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
		Status:      request.Status,
	}

	if err := db.Create(&scheduling).Error; err != nil {
		logger.Errorf("error create scheduling %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error create scheduling on database")
		return
	}
	sendSucess(ctx, "create scheduling", scheduling)
}
