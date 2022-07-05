package poker

import (
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
		return len(faces) >= 3
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
