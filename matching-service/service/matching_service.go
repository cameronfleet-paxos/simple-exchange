package service




/*
	POST
		/match (Order)
	logic = 
		get all orders from order service for the opposite symbol
		match bid/ask prices (exact only)
		create Match object

 */
type MatchService struct {

}