package handler

import (
	"fmt"
	"net/http"

	"github.com/approved-designs/simple-exchange/order-service/model"
	"github.com/approved-designs/simple-exchange/order-service/store"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderStore store.OrderStore
}

func (h OrderHandler) GetAll(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, h.OrderStore.GetAll())
}

func (h OrderHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, order := range h.OrderStore.GetAll() {
		if order.Id.String() == id {
			ctx.IndentedJSON(http.StatusOK, order)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "order not found"})
}

func (h OrderHandler) NewOrder(ctx *gin.Context) {
	var newOrder model.Order

	if err := ctx.BindJSON(&newOrder); err != nil {
		fmt.Println("Got err:", err)
		return
	}

	h.OrderStore.Upsert(newOrder)
	ctx.IndentedJSON(http.StatusCreated, newOrder)
}
