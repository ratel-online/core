package poker

import (
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/model"
	"github.com/ratel-online/core/util/arrays"
	"math"
	"sort"
)

func RunFastDistribute(dontShuffle bool, rules Rules) []model.Pokers {
	pokers := append(make(model.Pokers, 0), GetRunFastDontShuffleBase()...)
	for i := range pokers {
		pokers[i].Val = rules.Value(pokers[i].Key)
	}
	size := len(pokers)
	if dontShuffle {
		pokers.Shuffle(size, 3)
	} else {
		pokers.Shuffle(size, 1)
	}
	reserve := 0
	avgNum := (size - reserve) / 3
	pokersArr := make([]model.Pokers, 0)
	for i := 0; i < 3; i++ {
		pokerArr := make([]model.Poker, 0)
		pokersArr = append(pokersArr, append(pokerArr, pokers[i*avgNum:(i+1)*avgNum]...))
	}
	for i := range pokersArr {
		pokersArr[i].SortByValue()
	}
	return pokersArr
}

func RunFastParseFaces(pokers model.Pokers, rules Rules) []model.Faces {
	if len(pokers) == 0 {
		return nil
	}
	score := int64(0)
	stats := map[int]int{}
	group := map[int][]int{}
	counts := make([]int, 0)
	values := make([]int, 0)
	//單排記分
	for _, poker := range pokers {
		if poker.Key < 0 || poker.Key > 15 {
			return nil
		}
		poker.Val = rules.Value(poker.Key)
		score += int64(poker.Val)
		values = append(values, poker.Val)
		stats[poker.Val]++
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
	list := make([]model.Faces, 0)
	if 4 == len(pokers) && len(group[4]) == 1 {
		list = append(list, model.Faces{Values: values, Score: int64(group[4][0])*2 + int64(len(pokers)*1000), Type: consts.FacesBomb})
	} else if counts[0] == 1 {
		if len(group[counts[0]]) == 1 {
			list = append(list, model.Faces{Values: values, Score: score, Type: consts.FacesSingle})
		} else if rules.IsStraight(group[counts[0]], counts[0]) {
			list = append(list, model.Faces{Values: values, Score: score, Main: len(group[counts[0]]), Type: consts.FacesStraights})
		}
	} else if counts[0] == 2 && len(counts) == 1 {
		if len(group[counts[0]]) == 1 {
			list = append(list, model.Faces{Values: values, Score: score, Type: consts.FacesDouble})
		} else if rules.IsStraight(group[counts[0]], counts[0]) {
			list = append(list, model.Faces{Values: values, Score: score, Main: len(group[counts[0]]), Type: consts.FacesDoubles})
		}
	} else if counts[0] >= 3 {
		if len(counts) == 1 && len(group[counts[0]]) == 1 {
			if counts[0] == 3 {
				list = append(list, model.Faces{Values: values, Score: score, Type: consts.FacesTriple})
			} else {
				list = append(list, model.Faces{Values: values, Score: int64(group[counts[0]][0]*counts[0]) + int64(len(pokers)*1000), Type: consts.FacesBomb})
			}
		} else if len(counts) == 1 && rules.IsStraight(group[counts[0]], counts[0]) {
			if counts[0] == 3 {
				list = append(list, model.Faces{Values: values, Score: score, Main: len(group[counts[0]]), Type: consts.FacesThreeStraights})
			} else if counts[0] == 4 {
				values = make([]int, 0)
				for _, v := range group[counts[0]] {
					values = arrays.AppendN(values, v, 3)
				}
				for _, v := range group[counts[0]] {
					values = append(values, v)
				}
				list = append(list, model.Faces{Values: values, Score: score / 4 * 3, Main: len(group[counts[0]]), Extra: 1, Type: consts.FacesUnion3Straight})
			} else if counts[0] == 5 {
				values = make([]int, 0)
				for _, v := range group[counts[0]] {
					values = arrays.AppendN(values, v, 3)
				}
				for _, v := range group[counts[0]] {
					values = arrays.AppendN(values, v, 2)
				}
				list = append(list, model.Faces{Values: values, Score: score / 5 * 3, Main: len(group[counts[0]]), Extra: 1, Type: consts.FacesUnion3Straight})
			}
		} else if len(group[3]) > 0 || (len(group[4]) > 0 && len(values)-4 != 3) {
			//三帶二
			score := int64(0)
			main := 0
			extra := 0
			_type := consts.FacesUnion3c2
			if len(group[4]) == 1 && len(values)-3 == 2 {
				_type = consts.FacesUnion3c2s
				main = 3
				score = int64(group[4][0] * 3)
				list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
			} else if len(group[4]) == 1 && len(group[3]) >= 1 {
				temp := append([]int{group[4][0]})
				for k := range group[3] {
					temp = append(temp, group[3][k])
				}
				sort.Ints(temp)
				switch len(temp) - 1 {
				case 1:
					if math.Abs(float64(group[4][0]-group[3][0])) == 1 {
						_type = consts.FacesUnion3c2Cs
						main = 6
						extra = len(values) - main
						score = int64(group[4][0]*3 + group[3][0]*3)
						list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
					}
					break
				case 2:
					if temp[2]-temp[0] == 2 {
						_type = consts.FacesUnion3c2CsM
						main = 9
						extra = len(values) - main
						score = int64(group[4][0]*3 + group[3][0]*3 + group[3][1]*3)
						list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
					} else if temp[2]-temp[1] == 1 {
						_type = consts.FacesUnion3c2Cs
						main = 6
						extra = len(values) - main
						score = int64(temp[2]*3 + temp[1]*3)
						list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
					}
					break
				}
			} else if len(group[4]) == 2 && math.Abs(float64(group[4][0]-group[4][1])) <= 2 && len(group[3]) == 1 {
				if math.Abs(float64(group[4][0]-group[4][1])) == 2 {
					if math.Abs(float64(group[4][0]-group[3][0])) == 1 && math.Abs(float64(group[4][1]-group[3][0])) == 1 {
						_type = consts.FacesUnion3c2CsM
						main = 9
						extra = len(values) - main
						score += int64(group[4][0]*3 + group[4][1]*3 + group[3][0]*3)
						list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
					}
				} else {
					if math.Abs(float64(group[4][0]-group[3][0])) == 1 || math.Abs(float64(group[4][1]-group[3][0])) == 1 {
						_type = consts.FacesUnion3c2CsM
						main = 9
						extra = len(values) - main
						score += int64(group[4][0]*3 + group[4][1]*3 + group[3][0]*3)
						list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
					}
				}
			} else if len(group[4]) == 2 && len(values)-6 == 4 && math.Abs(float64(group[4][0]-group[4][1])) == 1 {
				_type = consts.FacesUnion3c2Cs
				main = 6
				extra = len(values) - main
				score += int64(group[4][0]*3 + group[4][1]*3)
				list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
			} else if len(group[4]) >= 2 {
				temp := make([]int, 0)
				for k := range group[4] {
					temp = append(temp, group[4][k])
				}
				for k := range group[3] {
					temp = append(temp, group[3][k])
				}
				sort.Ints(temp)
				for k := range temp {
					if k+2 >= len(temp) {
						break
					}
					if temp[k+2]-temp[k] == 2 {
						_type = consts.FacesUnion3c2CM
						if len(values)-9 == 6 {
							_type = consts.FacesUnion3c2CsM
						}
						main = 9
						extra = len(values) - main
						score += int64(temp[k]*3 + temp[k+1]*3 + temp[k+2]*3)
						list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
						break
					}
				}
			} else if len(group[3]) == 1 && len(values)-3 <= 2 {
				if len(values)-3 == 2 {
					_type = consts.FacesUnion3c2s
				}
				main = 3
				score = int64(group[3][0] * 3)
				list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
			} else if len(group[3]) == 2 && len(values)-6 <= 4 && math.Abs(float64(group[3][0]-group[3][1])) == 1 {
				_type = consts.FacesUnion3c2C
				if len(values)-6 == 4 {
					_type = consts.FacesUnion3c2Cs
				}
				main = 6
				extra = len(values) - main
				score += int64(group[3][0]*3 + group[3][1]*3)
				list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
			} else if len(group[3]) > 2 && len(values)-6 == 4 {
				temp := make([]int, 0)
				for k := range group[3] {
					temp = append(temp, group[3][k])
				}
				sort.Ints(temp)
				for k := range temp {
					if k+1 >= len(temp) {
						break
					}
					if temp[k+1]-temp[k] == 1 {
						_type = consts.FacesUnion3c2Cs
						main = 6
						extra = len(values) - main
						score += int64(temp[k]*3 + temp[k+1]*3)
						list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
						break
					}
				}
			} else if len(group[3]) == 3 && len(values)-9 <= 6 {
				_type = consts.FacesUnion3c2CM
				if len(values)-9 == 6 {
					_type = consts.FacesUnion3c2CsM
				}
				sort.Ints(group[3])
				if group[3][2]-group[3][0] == 2 {
					main = 9
					extra = len(values) - main
					score += int64(group[3][0]*3 + group[3][1]*3 + group[3][1]*3)
				} else {
					list = parseUnionOrStraight(group, rules)
					return list
				}
				list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
			} else if len(group[3]) > 3 && len(values)-9 <= 6 {
				//極端
				sort.Ints(group[3])
				for k := range group[3] {
					if k+2 >= len(group[3]) {
						break
					}
					if group[3][k+2]-group[3][k] == 2 {
						_type = consts.FacesUnion3c2CM
						if len(values)-9 == 6 {
							_type = consts.FacesUnion3c2CsM
						}
						main = 9
						extra = len(values) - main
						score += int64(group[3][k]*3 + group[3][k+1]*3 + group[3][k+2]*3)
						list = append(list, model.Faces{Values: values, Score: score, Main: main, Extra: extra, Type: consts.FacesType(_type)})
						break
					}
				}
			} else {
				list = parseUnionOrStraight(group, rules)
			}
		}
		if counts[0] == 4 && len(values)-4 <= 3 {
			_type := consts.FacesUnion4C3
			if len(values)-4 == 3 {
				_type = consts.FacesUnion4C3s
			}
			extra := len(values) - 4
			list = append(list, model.Faces{Values: values, Score: int64(group[counts[0]][0] * counts[0]), Main: len(group[counts[0]]), Extra: extra, Type: consts.FacesType(_type)})
		}
		if counts[0] == 4 && len(counts) == 1 && len(group[counts[0]]) == 2 {
			values = make([]int, 0)
			values = arrays.AppendN(values, group[counts[0]][1], counts[0])
			values = arrays.AppendN(values, group[counts[0]][0], counts[0])
			list = append(list, model.Faces{Values: values, Score: int64(group[counts[0]][1] * counts[0]), Main: 1, Extra: 1, Type: consts.FacesUnion4})
		}
	}
	return list
}

//RunFastComparativeFaces 计算能打的牌型
func RunFastComparativeFaces(lastPokers model.Faces, pokers model.Pokers, rules Rules) []model.Faces {
	accord := make([]model.Faces, 0)
	score := int64(0)
	stats := map[int]int{}
	valStats := map[int]int{}
	group := map[int][]int{}
	counts := make([]int, 0)
	values := make([]int, 0)
	keys := make([]int, 0)
	//單排記分
	for _, poker := range pokers {
		poker.Val = rules.Value(poker.Key)
		score += int64(poker.Val)
		values = append(values, poker.Val)
		keys = append(keys, poker.Key)
		stats[poker.Key]++
		valStats[rules.Value(poker.Key)]++
	}

	for v, c := range stats {
		group[c] = append(group[c], v)
	}

	for c := range group {
		counts = append(counts, c)
		sort.Ints(group[c])
	}
	sort.Ints(counts)

	switch lastPokers.Type {
	//炸彈處理
	case consts.FacesBomb:
		for k := len(group[4]) - 1; k >= 0; k-- {
			if rules.Value(group[4][k]) > lastPokers.Values[0] {
				accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesBomb, Keys: []int{group[4][k], group[4][k], group[4][k], group[4][k]}, Values: []int{rules.Value(group[4][k]), rules.Value(group[4][k]), rules.Value(group[4][k]), rules.Value(group[4][k])}})
			}
		}
		return accord
	//單排處理
	case consts.FacesSingle:
		// 倒序單排
		for k := len(keys) - 1; k >= 0; k-- {
			if rules.Value(keys[k]) > lastPokers.Values[0] {
				accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesSingle, Keys: []int{keys[k]}, Values: []int{rules.Value(keys[k])}})
			}
		}

		break
	//對子處理
	case consts.FacesDouble:
		//倒序對子
		for k := len(group[2]) - 1; k >= 0; k-- {
			if rules.Value(group[2][k]) > lastPokers.Values[0] {
				accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesDouble, Keys: []int{group[2][k], group[2][k]}, Values: []int{rules.Value(group[2][k]), rules.Value(group[2][k])}})
			}
		}
		// 倒序拆三張
		for k := len(group[3]) - 1; k >= 0; k-- {
			if rules.Value(group[3][k]) > lastPokers.Values[0] {
				accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesDouble, Keys: []int{group[3][k], group[3][k]}, Values: []int{rules.Value(group[3][k]), rules.Value(group[3][k])}})
			}
		}
		// 倒序拆炸彈
		for k := len(group[4]) - 1; k >= 0; k-- {
			if rules.Value(group[4][k]) > lastPokers.Values[0] {
				accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesDouble, Keys: []int{group[4][k], group[4][k]}, Values: []int{rules.Value(group[4][k]), rules.Value(group[4][k])}})
			}
		}
		break
	//三張處理
	case consts.FacesTriple, consts.FacesUnion3, consts.FacesUnion3c2, consts.FacesUnion3c2s:
		// 倒序三張
		for k := len(group[3]) - 1; k >= 0; k-- {
			//比分
			if int64(rules.Value(group[3][k])*3) > lastPokers.Score {
				key, val := RunFastMake(keys, []int{group[3][k], group[3][k], group[3][k]}, 2, rules)
				if len(keys)-3 <= 2 {
					accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion3c2, Keys: key, Values: val})
				} else {
					accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion3c2s, Keys: key, Values: val})
				}
			}
		}
		// 倒序拆炸彈
		for k := len(group[4]) - 1; k >= 0; k-- {
			if int64(rules.Value(group[4][k])*3) > lastPokers.Score {
				key, val := RunFastMake(keys, []int{group[4][k], group[4][k], group[4][k]}, 2, rules)
				if len(keys)-3 <= 2 {
					accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion3c2, Keys: key, Values: val})
				} else {
					accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion3c2s, Keys: key, Values: val})
				}
			}
		}
		break
	//二連三帶二處理
	case consts.FacesUnion3c2C, consts.FacesUnion3c2Cs:
		temp := make([]int, 0)
		for k := range group[4] {
			temp = append(temp, group[4][k])
		}
		for k := range group[3] {
			temp = append(temp, group[3][k])
		}
		sort.Ints(temp)
		for k := range temp {
			if k+1 >= len(temp) {
				break
			}
			if temp[k+1]-temp[k] == 1 {
				//比分數
				list := RunFastParseFaces(GetPokers(temp[k], temp[k], temp[k], temp[k+1], temp[k+1], temp[k+1]), rules)
				if len(list) >= 0 {
					if list[0].Score > lastPokers.Score {
						key, val := RunFastMake(keys, []int{temp[k], temp[k], temp[k], temp[k+1], temp[k+1], temp[k+1]}, 4, rules)
						if len(keys)-6 <= 4 {
							accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion3c2C, Keys: key, Values: val})
						} else {
							accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion3c2Cs, Keys: key, Values: val})
						}
					}
				}
			}
		}
		break
	//三連三帶二處理
	case consts.FacesUnion3c2CM, consts.FacesUnion3c2CsM:
		temp := make([]int, 0)
		for k := range group[4] {
			temp = append(temp, group[4][k])
		}
		for k := range group[3] {
			temp = append(temp, group[3][k])
		}
		sort.Ints(temp)
		for k := range temp {
			if k+2 >= len(temp) {
				break
			}
			if temp[k+2]-temp[k] == 2 {
				//比分數
				list := RunFastParseFaces(GetPokers(temp[k], temp[k], temp[k], temp[k+1], temp[k+1], temp[k+1], temp[k+2], temp[k+2], temp[k+2]), rules)
				if len(list) >= 0 {
					if list[0].Score > lastPokers.Score {
						key, val := RunFastMake(keys, []int{temp[k], temp[k], temp[k], temp[k+1], temp[k+1], temp[k+1], temp[k+2], temp[k+2], temp[k+2]}, 6, rules)
						if len(keys)-9 <= 6 {
							accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion3c2CM, Keys: key, Values: val})
						} else {
							accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion3c2CsM, Keys: key, Values: val})
						}
					}
				}
			}
		}
		break
	//四張處理
	case consts.FacesUnion4, consts.FacesUnion4C3, consts.FacesUnion4C3s:
		for k := len(group[4]) - 1; k >= 0; k-- {
			if int64(rules.Value(group[4][k]))*4 > lastPokers.Score {
				if len(keys) == 4 {
					break
				}
				key, val := RunFastMake(keys, []int{group[4][k], group[4][k], group[4][k], group[4][k]}, 3, rules)
				if len(keys)-4 <= 3 {
					accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion4C3, Keys: key, Values: val})
				} else {
					accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesUnion4C3s, Keys: key, Values: val})
				}
			}
		}
		break
		//連隊處理
	//連隊處理`
	case consts.FacesDoubles:
		if len(keys) >= len(lastPokers.Values) {
			temp := make([]int, 0)
			for k := range group[4] {
				temp = append(temp, group[4][k])
			}
			for k := range group[3] {
				temp = append(temp, group[3][k])
			}
			for k := range group[2] {
				temp = append(temp, group[2][k])
			}
			sort.Ints(temp)
			for k := range temp {
				if k+1 >= len(temp) {
					break
				}
				if temp[k+1]-temp[k] == 1 {
					//比分數
					list := RunFastParseFaces(GetPokers(temp[k], temp[k], temp[k+1], temp[k+1]), rules)
					if len(list) >= 0 {
						if list[0].Score > lastPokers.Score {
							accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesDoubles, Keys: []int{temp[k], temp[k], temp[k+1], temp[k+1]}, Values: []int{rules.Value(temp[k]), rules.Value(temp[k]), rules.Value(temp[k+1]), rules.Value(temp[k+1])}})
						}
					}
				}
			}
		}

		break

		//三連隊處理
	//三張隊處理
	case consts.FacesThreeStraights:
		if len(keys) >= len(lastPokers.Values) {
			temp := make([]int, 0)
			for k := range group[4] {
				temp = append(temp, group[4][k])
			}
			for k := range group[3] {
				temp = append(temp, group[3][k])
			}
			sort.Ints(temp)
			for k := range temp {
				if k+2 >= len(temp) {
					break
				}
				if temp[k+2]-temp[k] == 2 {
					//比分數
					list := RunFastParseFaces(GetPokers(temp[k], temp[k], temp[k], temp[k+1], temp[k+1], temp[k+1], temp[k+2], temp[k+2], temp[k+2]), rules)
					if len(list) >= 0 {
						if list[0].Score > lastPokers.Score {
							accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesThreeStraights, Keys: []int{temp[k], temp[k], temp[k], temp[k+1], temp[k+1], temp[k+1], temp[k+2], temp[k+2], temp[k+2]}, Values: []int{rules.Value(temp[k]), rules.Value(temp[k]), rules.Value(temp[k]), rules.Value(temp[k+1]), rules.Value(temp[k+1]), rules.Value(temp[k+1]), rules.Value(temp[k+2]), rules.Value(temp[k+2]), rules.Value(temp[k+2])}})
						}
					}
				}
			}
		}
		break
		//順子處理
	//順子處理
	case consts.FacesStraights:
		if len(keys) >= len(lastPokers.Values) {
			sort.Ints(lastPokers.Values)
			//不是封頂順子
			if lastPokers.Values[len(lastPokers.Values)-1] != 12 {
				ats := lastPokers.Values[0] + 1
				for {
					if ats+(len(lastPokers.Values)-1) <= 12 {
						ass := true
						val := make([]int, 0)
						key := make([]int, 0)
						tem := ats
						for i := 0; i < len(lastPokers.Values); i++ {
							_, ok := valStats[tem]
							if !ok {
								ass = false
							} else {
								val = append(val, tem)
								key = append(key, RunFastKey(tem))
							}
							tem++
						}
						if ass {
							accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesStraights, Keys: key, Values: val})
						}
						ats++
					} else {
						break
					}
				}

			}
		}
		break
	}
	// 倒序炸彈
	for k := len(group[4]) - 1; k >= 0; k-- {
		accord = append(accord, model.Faces{Score: lastPokers.Score + 1, Type: consts.FacesBomb, Keys: []int{group[4][k], group[4][k], group[4][k], group[4][k]}, Values: []int{rules.Value(group[4][k]), rules.Value(group[4][k]), rules.Value(group[4][k]), rules.Value(group[4][k])}})
	}
	return accord
}

//RunFastIsMax 是否為跑得快最大的牌 kkkk
func RunFastIsMax(faces model.Faces) bool {
	if len(faces.Keys) != 4 {
		return false
	}
	return faces.Keys[0] == 13 && faces.Keys[1] == 13 && faces.Keys[2] == 13 && faces.Keys[3] == 13
}

//RunFastKey 跑得快值返回key
func RunFastKey(val int) int {
	if val == 12 {
		return 1
	} else if val == 13 {
		return 2
	}
	return val + 2
}

//RunFastMake 跑得快带牌算法
func RunFastMake(keys []int, exclude []int, count int, rules Rules) ([]int, []int) {
	values := make([]int, 0)
	for _, key := range exclude {
		j := 0
		for _, val := range keys {
			if val != key {
				keys[j] = val
				j++
			}
		}
		keys = keys[:j]
	}
	for _, key := range keys {
		values = append(values, rules.Value(key))
	}
	sort.Ints(values)
	key := make([]int, 0)
	val := make([]int, 0)
	if len(keys) >= count {
		for i := 0; i < count; i++ {
			key = append(key, RunFastKey(values[i]))
			val = append(val, values[i])
		}
	} else {
		for i := 0; i < len(keys); i++ {
			key = append(key, RunFastKey(values[i]))
			val = append(val, values[i])
		}
	}
	for _, k := range exclude {
		key = append(key, k)
		val = append(val, rules.Value(k))
	}
	return key, val
}
