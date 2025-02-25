package poker

import (
	"github.com/ratel-online/core/model"
	"github.com/ratel-online/core/pkg/holdem"
)

func ParseTexasFaces(hand, board model.Pokers) (*model.TexasFaces, error) {
	// prevent slice append slice from modifying the original slice
	cards := make(model.Pokers, 0)
	cards = append(cards, hand...)
	cards = append(cards, board...)

	holdemCards := [7]holdem.Card{}
	for i, c := range cards {
		val := c.Key
		if val == 1 {
			val = 14
		}
		val <<= 4
		switch c.Suit {
		case model.Spade:
			val |= 1
		case model.Club:
			val |= 2
		case model.Heart:
			val |= 3
		case model.Diamond:
			val |= 4
		}
		holdemCards[i] = holdem.Card(val)
	}

	score, _ := holdem.HighestHandValue(holdemCards)
	return &model.TexasFaces{
		Type:  model.TexasFacesType(score>>20 + 1),
		Score: int64(score),
	}, nil
}
