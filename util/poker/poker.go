package poker

import (
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/model"
	"github.com/ratel-online/core/util/math"
	"github.com/ratel-online/core/util/strings"
	"sort"
)

type Rules interface {
	Score(key int) int
	IsStraight(faces []int, count int) bool
}

var base = make(model.Pokers, 0)

func init() {
	for k := 1; k <= 13; k++ {
		for t := 1; t <= 4; t++ {
			base = append(base, model.Poker{Key: k, Val: 0, Type: t, Desc: desc(k)})
		}
	}
	for k := 16; k <= 17; k++ {
		base = append(base, model.Poker{Key: k, Val: 0, Type: 1, Desc: desc(k)})
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
	sets := number / 3
	if number%3 > 0 {
		sets++
	}
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

func SetValue(pokers model.Pokers, kv map[int]int) {
	for i := range pokers {
		pokers[i].Val = kv[pokers[i].Key]
	}
}

func ParseFaces(pokers model.Pokers, rules Rules) *model.Faces {
	sc, xc, score := 0, 0, int64(0)
	stat := map[int]int{}
	group := map[int][]int{}
	keys := make([]int, 0)
	counts := make([]int, 0)
	for _, poker := range pokers {
		poker.Val = rules.Score(poker.Key)
		stat[poker.Val]++
		score += int64(poker.Val)
		if poker.Key == 14 {
			sc++
		} else if poker.Key == 15 {
			xc++
		}
	}
	for v, c := range stat {
		group[c] = append(group[c], v)
		counts = append(counts, c)
	}
	for c := range group {
		keys = append(keys, c)
		sort.Ints(group[c])
	}
	sort.Ints(keys)
	sort.Ints(counts)
	for i := 0; i < len(keys)/2; i++ {
		keys[i], keys[len(keys)-i-1] = keys[len(keys)-i-1], keys[i]
	}
	faces := &model.Faces{Counts: counts, Score: score}
	if len(keys) == 0 {
		return faces.SetType(consts.FacesInvalid)
	}
	if sc+xc == len(pokers) && sc+xc > 1 {
		return faces.SetScore(int64(sc*14+xc*15)*2 + int64(len(pokers)*2*1000)).SetType(consts.FacesBomb)
	}
	if keys[0] == 1 {
		if len(group[1]) == 1 {
			return faces.SetType(consts.FacesSingle)
		} else if rules.IsStraight(group[1], keys[0]) {
			return faces.SetType(consts.FacesStraight)
		}
	} else if keys[0] == 2 {
		if len(keys) == 1 {
			if len(group[2]) == 1 {
				return faces.SetType(consts.FacesDouble)
			} else if rules.IsStraight(group[2], keys[0]) {
				return faces.SetType(consts.FacesStraight)
			}
		}
	} else if keys[0] >= 3 {
		if len(keys) == 1 {
			if len(group[keys[0]]) == 1 {
				if keys[0] == 3 {
					return faces.SetType(consts.FacesTriple)
				} else {
					return faces.SetScore(int64(group[keys[0]][0]*keys[0]) + int64(len(pokers)*1000)).SetType(consts.FacesBomb)
				}
			} else if rules.IsStraight(group[keys[0]], keys[0]) {
				return faces.SetType(consts.FacesStraight)
			} else if len(group[keys[0]]) == 2 && keys[0] == 4 {
				return faces.SetType(consts.FacesUnion)
			}
		} else if len(keys) == 2 && keys[1] <= 2 {
			ml := len(group[keys[0]])
			al := len(group[keys[1]])

			if (keys[0] == 3 && (ml == al || (keys[1] == 2 && ml == 2*al))) || (al%ml == 0 && (al/ml == 2 || (keys[1] == 2 && ml/al == 1))) {
				if len(group[keys[0]]) == 1 {
					return faces.SetScore(int64(math.Sum(group[keys[0]]...) * keys[0])).SetType(consts.FacesUnion)
				} else if rules.IsStraight(group[keys[0]], keys[0]) {
					return faces.SetScore(int64(math.Sum(group[keys[0]]...) * keys[0])).SetType(consts.FacesUnionStraight)
				}
			}
		}
	}
	return faces.SetType(consts.FacesInvalid)
}
