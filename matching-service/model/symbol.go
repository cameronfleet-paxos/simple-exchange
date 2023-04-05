package model

import "fmt"

type Symbol struct {
	BidAsset Asset
	AskAsset Asset
}

func (s Symbol) String() string {
	return fmt.Sprintf("%s/%s", s.BidAsset, s.AskAsset)
}

func (s Symbol) Opposite() Symbol {
	return Symbol{BidAsset: s.AskAsset, AskAsset: s.BidAsset}
}
