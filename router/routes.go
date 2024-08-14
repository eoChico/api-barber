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
		v1.GET("/scheduling", handler.ShowScheduling)
		v1.DELETE("/scheduling", handler.DeleteScheduling)
		v1.POST("/barber", handler.CreateBarber)
		//v1.GET("/barbers", handler.ListBarbers)
	}
}
