package main

import (
	"github.com/approved-designs/simple-exchange/matching-service/external"
	"github.com/approved-designs/simple-exchange/matching-service/matcher"
	"github.com/approved-designs/simple-exchange/matching-service/service"
)

func main() {
	orderMatcher := &matcher.OrderMatcher{OrderService: &external.HttpOrderService{}}
	matchingService := service.MatchService{OrderMatcher: orderMatcher}
	matchingService.Start()
}
