package service

import (
	"github.com/gin-gonic/gin"
)

type OrderMatcher interface {
	MatchOrder(ctx *gin.Context)
}

type MatchService struct {
	OrderMatcher OrderMatcher
}

func (m *MatchService) Start() {
	router := gin.Default()
	router.POST("/match", m.OrderMatcher.MatchOrder)
	router.Run("localhost:8080")
}
