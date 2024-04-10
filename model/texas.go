package model

type TexasFacesType int

const (
	_                           TexasFacesType = iota
	TexasFacesTypeHigh                         // 高牌
	TexasFacesTypeOnePair                      // 一对
	TexasFacesTypeTwoPairs                     // 两对
	TexasFacesTypeThreeOfAKind                 // 三条
	TexasFacesTypeStraight                     // 顺子
	TexasFacesTypeFlush                        // 同花
	TexasFacesTypeFullHouse                    // 葫芦
	TexasFacesTypeFourOfAKind                  // 四条
	TexasFacesTypeStraightFlush                // 同花顺
	TexasFacesTypeRoyalFlush                   // 皇家同花顺
)

func (t TexasFacesType) String() string {
	switch t {
	case TexasFacesTypeHigh:
		return "高牌(High)"
	case TexasFacesTypeOnePair:
		return "一对(One Pair)"
	case TexasFacesTypeTwoPairs:
		return "两对(Two Pairs)"
	case TexasFacesTypeThreeOfAKind:
		return "三条(Three of a Kind)"
	case TexasFacesTypeStraight:
		return "顺子(Straight)"
	case TexasFacesTypeFlush:
		return "同花(Flush)"
	case TexasFacesTypeFullHouse:
		return "葫芦(Full House)"
	case TexasFacesTypeFourOfAKind:
		return "四条(Four of a Kind)"
	case TexasFacesTypeStraightFlush:
		return "同花顺(Straight Flush)"
	case TexasFacesTypeRoyalFlush:
		return "皇家同花顺(Royal Flush)"
	}
	return "Unknown"
}

type TexasFaces struct {
	Score int64          `json:"score"`
	Type  TexasFacesType `json:"type"`
}
