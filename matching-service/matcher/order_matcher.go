package matcher

import (
	"fmt"
	"net/http"

	"github.com/approved-designs/simple-exchange/matching-service/external"
	"github.com/approved-designs/simple-exchange/matching-service/model"
	"github.com/gin-gonic/gin"
)

type OrderMatcher struct {
	OrderService external.OrderService
}

func (o *OrderMatcher) MatchOrder(c *gin.Context) {
	var unmatchedOrder model.UnmatchedOrder

	if err := c.BindJSON(&unmatchedOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	orders, err := o.OrderService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	fmt.Println("Attempting to match order:", unmatchedOrder, "with orders:", orders)
	for _, order := range orders {
		if order.Symbol == unmatchedOrder.Symbol.Opposite() {
			c.JSON(http.StatusOK, gin.H{"message": "order matched with: " + order.Id.String()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "order unmatched"})
}
