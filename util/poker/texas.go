package poker

import (
	"github.com/ratel-online/core/model"
	handx "github.com/ratel-online/core/pkg/hand"
)

func ParseTexasFaces(hand, board model.Pokers) (*model.TexasFaces, error) {
	h := handx.GetHand()
	h.Init()

	for _, c := range append(hand, board...) {
		val := c.Key
		if val == 1 {
			val = 14
		}
		val -= 2
		err := h.SetCard(&handx.Card{
			Suit:  int(c.Suit),
			Value: val,
		})
		if err != nil {
			return nil, err
		}
	}
	err := h.AnalyseHand()
	if err != nil {
		return nil, err
	}

	return &model.TexasFaces{
		Type:  model.TexasFacesType(h.Level),
		Score: int64(h.FinalValue),
	}, nil
}
