package poker

import (
	"github.com/ratel-online/core/model"
	"github.com/ratel-online/core/util/strings"
)

var base = make(model.Pokers, 0)

func init() {
	for k := 1; k <= 13; k++ {
		for t := 1; t <= 4; t++ {
			base = append(base, model.Poker{
				Key:  k,
				Val:  0,
				Type: t,
				Desc: desc(k),
			})
		}
	}
}

func desc(k int) string {
	switch k {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "S"
	case 15:
		return "X"
	default:
		return strings.String(k)
	}
}

func Distribute(number int, reserved bool) []model.Pokers {
	sets := number/3 + 1
	pokers := make(model.Pokers, 0)
	for i := 0; i < sets; i++ {
		pokers = append(pokers, base...)
	}
	pokers.Shuffle()
	size := len(pokers)
	reserve := 0
	if reserved {
		if size%number == 0 {
			reserve = number * sets
		} else {
			reserve = number + size%number
		}
	} else {
		reserve = size % number
	}
	avgNum := (size - reserve) / number
	pokersArr := make([]model.Pokers, 0)
	for i := 0; i < number; i++ {
		pokersArr = append(pokersArr, pokers[i*avgNum:(i+1)*avgNum])
	}
	if reserve > 0 {
		pokersArr = append(pokersArr, pokers[size-reserve:])
	}
	return pokersArr
}
