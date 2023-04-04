package service

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	NewOrder(ctx *gin.Context)
}

type OrderService struct {
	OrderHandler OrderHandler
}

func (s *OrderService) Start() {
	router := gin.Default()
	router.GET("/orders", s.OrderHandler.GetAll)
	router.GET("/order/:id", s.OrderHandler.GetById)
	router.POST("/order", s.OrderHandler.NewOrder)
	router.Run("localhost:8080")
}
