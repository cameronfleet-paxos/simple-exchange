package model

import (
	"time"

	"github.com/google/uuid"
)

type MatchId uuid.UUID

type Match struct {
	MatchId MatchId
	LeftOrderId OrderId
	RightOrderId OrderId
	CreatedAt time.Time
}