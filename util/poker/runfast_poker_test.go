package poker

import (
	"fmt"
	"github.com/ratel-online/core/model"
	"testing"
)

var runFastRules = _runFastRules{}

type _runFastRules struct {
}

func (d _runFastRules) Value(key int) int {
	if key == 1 {
		return 12
	} else if key == 2 {
		return 13
	} else if key > 13 {
		return key
	}
	return key - 2
}

func (d _runFastRules) IsStraight(faces []int, count int) bool {
	if faces[len(faces)-1]-faces[0] != len(faces)-1 {
		return false
	}
	if faces[len(faces)-1] > 12 {
		return false
	}
	if count == 1 {
		return len(faces) >= 5
	} else if count == 2 {
		return len(faces) >= 2
	} else if count > 2 {
		return len(faces) >= 2
	}
	return false
}

func (d _runFastRules) StraightBoundary() (int, int) {
	return 1, 12
}

func (d _runFastRules) Reserved() bool {
	return false
}

//跑得快分牌测试
func TestRunFastDistribute(t *testing.T) {
	for i := 0; i < 100; i++ {
		pokersArr := RunFastDistribute(false, runFastRules)
		for _, pokers := range pokersArr {
			pokers.SortByValue()
			t.Log(pokers.String())
		}
		t.Log("\n")
	}

	//不洗牌模式
	//pokersArr := RunFastDistribute(true, runFastRules)
	//for _, pokers := range pokersArr {
	//	pokers.SortByValue()
	//	t.Log(pokers.String())
	//}
}

func TestRunFastParseFacesScore(t *testing.T) {
	testCases := []parseFacesCase{
		//10 9 5 5 5 4 4 4
		//
		{pokers: getPokers(6, 6, 7, 7)},
		{pokers: getPokers(10, 10, 10, 11, 11, 11, 8, 8, 8, 5)},
		{pokers: getPokers(7, 7, 7, 8, 8, 8, 3, 3, 3, 7)},
		{pokers: getPokers(7, 7, 7, 8, 8, 8, 3, 3, 3, 3)},

		//順子 or 連隊
		{pokers: getPokers(8, 8, 9, 9)},
		{pokers: getPokers(8, 8, 9, 9, 10, 10)},
		{pokers: getPokers(8, 8, 9, 9, 10, 10, 11, 11)},
		{pokers: getPokers(6, 7, 8, 9, 10)},
		{pokers: getPokers(7, 8, 9, 10, 11)},
		{pokers: getPokers(4, 5, 6, 7, 8, 9, 10)},
		{pokers: getPokers(5, 6, 7, 8, 9, 10, 11)},
		//三張
		{pokers: getPokers(6, 6, 6)},
		//三張連隊
		{pokers: getPokers(6, 6, 6, 7, 7, 7)},
		{pokers: getPokers(6, 6, 6, 7, 7, 7, 8, 8, 8)},
		{pokers: getPokers(6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9)},
		{pokers: getPokers(6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 10)},
		//三帶2
		{pokers: getPokers(6, 6, 6, 7)},
		{pokers: getPokers(6, 6, 6, 6, 7)},
		//四帶三
		{pokers: getPokers(6, 6, 6, 6, 7, 7)},
		{pokers: getPokers(6, 6, 6, 6, 7, 7, 7)},
		{pokers: getPokers(6, 6, 6, 6, 7, 7, 7, 4, 5, 5)},
		{pokers: getPokers(5, 6, 7, 8, 9, 10, 11, 12, 13)},
		{pokers: getPokers(6, 6, 6, 6, 7, 7, 7, 7, 5, 5, 5, 4, 9, 8, 3)},
		{pokers: getPokers(6, 6, 6, 7, 7, 7, 8, 8, 8, 6, 7, 8, 3, 3, 3)},
		{pokers: getPokers(6, 6, 6, 7, 7, 7, 8, 8, 8, 5, 5, 5, 3, 3, 3)},
		{pokers: getPokers(6, 6, 6, 7, 7, 7, 5, 5, 5, 7, 5, 10, 10, 10, 10)},
		{pokers: getPokers(9, 9, 9, 9, 5)},
		{pokers: getPokers(3, 3, 3, 3)},
		{pokers: getPokers(2, 2, 2, 2)},
		//連隊
		{pokers: getPokers(11, 11, 11, 11)},
		{pokers: getPokers(13, 13, 13, 13)},
	}
	preScore := int64(-1)
	for _, testCase := range testCases {
		list := RunFastParseFaces(testCase.pokers, runFastRules)
		if len(list) > 0 {
			faces := list[0]
			if faces.Score < preScore {
				t.Error(fmt.Sprintf("err score, pre %v should lt %v", preScore, faces.Score))
			}
			preScore = faces.Score
			t.Log(testCase.pokers.String(), faces.Score, faces.Type)
		}
	}
}

func TestRunFastComparativeFaces(t *testing.T) {
	list := RunFastComparativeFaces(model.Faces{Keys: []int{6, 6, 7, 7}, Values: []int{4, 4, 5, 5}, Score: int64(18), Main: 1, Type: 18}, getPokers(7, 7, 6, 6), runFastRules)
	if len(list) > 0 {
		faces := list[0]
		t.Log(faces.Values, faces.Score, faces.Type)
	}
}
