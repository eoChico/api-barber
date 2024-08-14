package handler

import (
	"net/http"

	"github.com/eoChico/api-barber/models"
	"github.com/gin-gonic/gin"
)

func CreateBarber(ctx *gin.Context) {
	request := CreateBarberRequest{}
	ctx.BindJSON(&request)

	barber := models.Barber{
		Name:  request.Name,
		Email: request.Email,
		Phone: request.Phone,
	}

	if err := db.Create(&barber).Error; err != nil {
		logger.Errorf("error create barber %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error create barber on database")
		return
	}
	sendSucess(ctx, "create barber", barber)
}
