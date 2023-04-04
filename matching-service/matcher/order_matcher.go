package matcher

import (
	"fmt"

	"github.com/approved-designs/simple-exchange/matching-service/model"
	"github.com/gin-gonic/gin"
)

type OrderMatcher struct {

}

func (o OrderMatcher) MatchOrder(ctx *gin.Context) {
	var unmatchedOrder model.UnmatchedOrder

	if err := ctx.BindJSON(&unmatchedOrder); err != nil {
		fmt.Println("Got err:", err)
		return
	}

	// OrderService.getAll()
	// Match
	// Create matched orders
	// Update orders on order service
}