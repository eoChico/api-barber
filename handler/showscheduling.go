package handler

import (
	"net/http"
	"time"

	"github.com/eoChico/api-barber/models"
	"github.com/gin-gonic/gin"
)

func ShowScheduling(ctx *gin.Context) {
	var schedulingsinDay []CreateSchedulingRequest
	request := ShowSchedulingsRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	day := request.Day
	startOfDay := day
	endOfDay := day.AddDate(0, 0, 1).Add(-time.Nanosecond)

	var schedulings []models.Scheduling
	db.Where("start_time BETWEEN ? AND ?", startOfDay, endOfDay).Find(&schedulings)

	for _, scheduling := range schedulings {
		//fmt.Println(scheduling.Description, scheduling.StartTime)
		appoiment := CreateSchedulingRequest{
			StartTime:   scheduling.StartTime,
			Description: scheduling.Description,
			EndTime:     scheduling.EndTime,
			Status:      scheduling.Status,
		}
		schedulingsinDay = append(schedulingsinDay, appoiment)
	}
	sendSucess(ctx, "learn scheduling", schedulingsinDay)
}
