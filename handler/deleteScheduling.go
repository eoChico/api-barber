package handler

import (
	"fmt"
	"net/http"

	"github.com/eoChico/api-barber/models"
	"github.com/gin-gonic/gin"
)

func DeleteScheduling(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsisRequired("id", "queryParameter").Error())
		return
	}
	scheduling := models.Scheduling{}
	//find opening
	if err := db.First(&scheduling, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("scheduling with id:%s not found", id))
		return
	}
	// delete opening
	if err := db.Delete(&scheduling).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting scheduling with id: %s", id))
		return
	}
	sendSucess(ctx, "delete scheduling", scheduling)
}
