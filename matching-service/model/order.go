package model

import "github.com/google/uuid"

type OrderId uuid.UUID

type Order struct {
	Id       OrderId
	Symbol   Symbol
	BidPrice uint64
	AskPrice uint64
}

type UnmatchedOrder Order

type MatchedOrder struct {
	Order
	MatchId MatchId
}