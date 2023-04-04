package model

import "github.com/google/uuid"

type Order struct {
	Id       uuid.UUID
	Symbol   Symbol
	BidPrice uint64
	AskPrice uint64
	MatchId  uuid.UUID
}
