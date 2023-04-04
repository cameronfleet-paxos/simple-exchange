package model

import "fmt"

type Symbol struct {
	BidAsset Asset
	AskAsset Asset
}

func (s Symbol) String() string {
	return fmt.Sprintf("%s/%s", s.BidAsset, s.AskAsset)
}
