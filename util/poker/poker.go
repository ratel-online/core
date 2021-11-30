package poker

import (
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/model"
	"github.com/ratel-online/core/util/strings"
	"sort"
)

type Rules interface {
	Score(key int) int
	IsStraight(faces []int, count int) bool
	StraightBoundary() (int, int)
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
	stats := map[int]int{}
	group := map[int][]int{}
	counts := make([]int, 0)
	keys := make([]int, 0)
	for _, poker := range pokers {
		poker.Val = rules.Score(poker.Key)
		keys = append(keys, poker.Key)
		stats[poker.Val]++
		score += int64(poker.Val)
		if poker.Key == 14 {
			sc++
		} else if poker.Key == 15 {
			xc++
		}
	}
	for v, c := range stats {
		group[c] = append(group[c], v)
	}
	for c := range group {
		counts = append(counts, c)
		sort.Ints(group[c])
	}
	sort.Ints(counts)
	for i := 0; i < len(counts)/2; i++ {
		counts[i], counts[len(counts)-i-1] = counts[len(counts)-i-1], counts[i]
	}
	faces := &model.Faces{Keys: keys, Score: score}
	if len(counts) == 0 {
		return faces.SetType(consts.FacesInvalid)
	}
	if sc+xc == len(pokers) && sc+xc > 1 {
		return faces.SetScore(int64(sc*14+xc*15)*2 + int64(len(pokers)*2*1000)).SetType(consts.FacesBomb)
	}
	if counts[0] == 1 {
		if len(group[counts[0]]) == 1 {
			return faces.SetType(consts.FacesSingle)
		} else if rules.IsStraight(group[counts[0]], counts[0]) {
			return faces.SetMain(len(group[counts[0]])).SetType(consts.FacesStraight)
		}
	} else if counts[0] == 2 && len(counts) == 1 {
		if len(group[counts[0]]) == 1 {
			return faces.SetType(consts.FacesDouble)
		} else if rules.IsStraight(group[counts[0]], counts[0]) {
			return faces.SetMain(len(group[counts[0]])).SetType(consts.FacesStraight)
		}
	} else if counts[0] >= 3 {
		if len(counts) == 1 && len(group[counts[0]]) == 1 {
			if counts[0] == 3 {
				return faces.SetType(consts.FacesTriple)
			} else {
				return faces.SetScore(int64(group[counts[0]][0]*counts[0]) + int64(len(pokers)*1000)).SetType(consts.FacesBomb)
			}
		} else if len(counts) == 1 && rules.IsStraight(group[counts[0]], counts[0]) {
			if counts[0] == 3 {
				return faces.SetMain(len(group[counts[0]])).SetType(consts.FacesStraight)
			} else if counts[0] == 4 {
				keys = make([]int, 0)
				for _, v := range group[counts[0]] {
					for i := 0; i < 3; i++ {
						keys = append(keys, v)
					}
				}
				for _, v := range group[counts[0]] {
					for i := 0; i < 1; i++ {
						keys = append(keys, v)
					}
				}
				return faces.SetMain(len(group[counts[0]])).SetExtra(1).SetKeys(keys).SetScore(score / 4 * 3).SetType(consts.FacesUnionStraight)
			} else if counts[0] == 5 {
				keys = make([]int, 0)
				for _, v := range group[counts[0]] {
					for i := 0; i < 3; i++ {
						keys = append(keys, v)
					}
				}
				for _, v := range group[counts[0]] {
					for i := 0; i < 2; i++ {
						keys = append(keys, v)
					}
				}
				return faces.SetMain(len(group[counts[0]])).SetExtra(1).SetKeys(keys).SetScore(score / 5 * 3).SetType(consts.FacesUnionStraight)
			}
		} else if len(group[3]) > 0 {
			faces, err := parseUnionOrStraight(group, rules)
			if err == nil {
				return faces
			}
		} else if counts[0] == 4 && len(counts) == 2 && counts[1] <= 2 {
			if len(group[counts[0]]) == 1 && ((counts[1] == 2 && len(group[counts[1]]) <= 2) || (counts[1] == 1 && len(group[counts[1]]) == 2)) {
				return faces.SetMain(len(group[counts[0]])).SetExtra(counts[1]).SetScore(int64(group[counts[0]][0] * counts[0])).SetType(consts.FacesUnion)
			}
		}
	}
	return faces.SetType(consts.FacesInvalid)
}

func parseUnionOrStraight(group map[int][]int, rules Rules) (*model.Faces, error) {
	extras := map[int]int{}
	mains := make([]int, 0)
	for k, arr := range group {
		for _, v := range arr {
			if k > 3 {
				extras[v] += k - 3
				mains = append(mains, v)
			} else if k == 3 {
				mains = append(mains, v)
			} else if k < 3 {
				extras[v] += k
			}
		}
	}
	sort.Ints(mains)
	valid := map[int]int{}
	sta, pre := mains[0], mains[0]
	for i := 1; i < len(mains); i++ {
		if mains[i] > pre+1 {
			valid[sta] = pre
			sta = mains[i]
		}
		pre = mains[i]
	}
	valid[sta] = mains[len(mains)-1]

	target := 0
	for k, v := range valid {
		if target == 0 {
			target = k
			continue
		}
		if v-k >= valid[target]-target && k > target {
			target = k
		}
	}

	for _, v := range mains {
		if v < target && v > valid[target] {
			extras[v] += 3
		}
	}

	ml, mr := target, valid[target]
	ll, lr := rules.StraightBoundary()
	main, extra, single, double := mr-ml+1, 0, 0, 0
	for _, v := range extras {
		single += v
		double += v / 2
	}

	extra = 1
	access, vl, vr := isValidUnionStraight(main, single, ml, mr, ll, lr)
	if !access && single-double*2 == 0 {
		extra = 2
		access, vl, vr = isValidUnionStraight(main, double, ml, mr, ll, lr)
	}
	if !access {
		return nil, consts.ErrsInvalidFaces
	}
	if vl > ml {
		for i := ll; i < vl; i++ {
			extras[i] += 3
		}
	}
	if vr < mr {
		for i := vr + 1; i <= mr; i++ {
			extras[i] += 3
		}
	}
	keys := make([]int, 0)
	arr := make([]int, 0)
	score := 0
	main = vr - vl + 1
	for i := vl; i <= vr; i++ {
		arr = append(arr, i)
		for j := 0; j < 3; j++ {
			keys = append(keys, i)
		}
		score += 3 * i
	}
	for k, v := range extras {
		for j := 0; j < v; j++ {
			keys = append(keys, k)
		}
	}
	faces := &model.Faces{
		Keys:  keys,
		Main:  main,
		Extra: extra,
		Score: int64(score),
	}
	if main == 1 {
		return faces.SetType(consts.FacesUnion), nil
	}
	return faces.SetType(consts.FacesUnionStraight), nil
}

func isValidUnionStraight(main, extras int, ml, mr, ll, lr int) (bool, int, int) {
	if main > extras && (main-extras)%3 == 0 {
		diff := (main - extras) / 3
		for i := 0; i < diff; i++ {
			if mr > lr {
				mr--
			} else {
				ml++
			}
		}
		main = mr - ml
		extras = main
	}
	return main == extras && ml >= ll && mr <= lr, ml, mr
}
