package routes

import (
	"restapigo/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders", handlers.GetOrders)
	router.PUT("/orders/:id", handlers.UpdateOrder)
	router.DELETE("/orders/:id", handlers.DeleteOrder)
}
