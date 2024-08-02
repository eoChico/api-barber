package router

import (
	"github.com/eoChico/api-barber/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	handler.InitHandler()
	v1 := r.Group("/api")
	{
		v1.POST("/scheduling", handler.CreateScheduling)
	}
}
