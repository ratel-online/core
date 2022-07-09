package poker

import (
	"fmt"
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/model"
	"testing"
)

var defaultRules = _defaultRules{}

type _defaultRules struct {
}

func (d _defaultRules) Value(key int) int {
	if key == 1 {
		return 12
	} else if key == 2 {
		return 13
	} else if key > 13 {
		return key
	}
	return key - 2
}

func (d _defaultRules) IsStraight(faces []int, count int) bool {
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

func (d _defaultRules) StraightBoundary() (int, int) {
	return 1, 12
}

func (d _defaultRules) Reserved() bool {
	return true
}

func getPokers(keys ...int) model.Pokers {
	pokers := make(model.Pokers, 0)
	for _, k := range keys {
		pokers = append(pokers, model.Poker{
			Key:  k,
			Desc: desc[k],
		})
	}
	return pokers
}

func TestDistribute(t *testing.T) {
	//pokersArr, _ := Distribute(3, defaultRules)
	//for _, pokers := range pokersArr {
	//	pokers.SortByKey()
	//	t.Log(pokers.String())
	//}
	//pokersArr, _ = Distribute(5, defaultRules)
	//for _, pokers := range pokersArr {
	//	pokers.SortByKey()
	//	t.Log(pokers.String())
	//}
	//pokersArr, _ = Distribute(7, defaultRules)
	//for _, pokers := range pokersArr {
	//	pokers.SortByKey()
	//	t.Log(pokers.String())
	//}
	pokersArr, _ := Distribute(3, true, defaultRules)
	for _, pokers := range pokersArr {
		pokers.SortByKey()
		t.Log(pokers.String())
	}
}

type parseFacesCase struct {
	pokers     model.Pokers
	actualType consts.FacesType
}

func testParseFaces(t *testing.T, pokers model.Pokers, expected consts.FacesType) {
	list := ParseFaces(pokers, defaultRules)
	if expected == consts.FacesInvalid && len(list) > 0 {
		t.Error("err at", pokers.String(), "expected invalid", "actual ", list[0].Values)
		return
	}
	access := false
	t.Log(pokers.String())
	for _, faces := range list {
		if expected == faces.Type {
			t.Log("\t\t->", faces.Keys)
			access = true
		}
	}
	if !access && expected != consts.FacesInvalid {
		t.Error("err at", pokers.String(), "expected ", expected, "actual ", list[0].Values)
	}
}

func TestParseFaces(t *testing.T) {
	testCases := []parseFacesCase{
		{getPokers(3), consts.FacesSingle},
		{getPokers(15), consts.FacesSingle},
		{getPokers(3, 3), consts.FacesDouble},
		{getPokers(3, 3, 3), consts.FacesTriple},
		{getPokers(3, 3, 3, 4), consts.FacesUnion3},
		{getPokers(3, 3, 3, 4, 4), consts.FacesUnion3},
		{getPokers(3, 3, 3, 4, 4, 5), consts.FacesInvalid},
		{getPokers(3, 3, 3, 3, 4), consts.FacesInvalid},
		{getPokers(3, 3, 3, 3, 4, 5), consts.FacesUnion4},
		{getPokers(3, 3, 3, 3, 3, 4, 5), consts.FacesInvalid},
		{getPokers(3, 4, 5, 6, 7), consts.FacesStraight},
		{getPokers(3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1), consts.FacesStraight},
		{getPokers(2, 3, 4, 5, 6, 7), consts.FacesInvalid},
		{getPokers(10, 11, 12, 13, 1, 2), consts.FacesInvalid},
		{getPokers(3, 3, 4, 4, 5, 5), consts.FacesStraight},
		{getPokers(3, 3, 3, 4, 4, 4), consts.FacesStraight},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 4), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 3, 3, 4, 4, 4, 4, 4), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 6), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5, 6, 6), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 4, 5, 6), consts.FacesInvalid},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6), consts.FacesInvalid},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 4, 5, 6, 7, 8), consts.FacesInvalid},
		{getPokers(3, 3, 3, 3), consts.FacesBomb},
		{getPokers(3, 3, 3, 3, 3), consts.FacesBomb},
		{getPokers(14, 15), consts.FacesBomb},
		{getPokers(14, 15, 15), consts.FacesBomb},
		{getPokers(14, 14, 15, 15), consts.FacesBomb},
		{getPokers(14, 14, 14, 15, 15, 16), consts.FacesInvalid},
		{getPokers(4, 4, 4, 4, 6, 6, 6, 6), consts.FacesInvalid},
		{getPokers(4, 4, 4, 16), consts.FacesInvalid},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5, 5, 7, 7, 7), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 7, 7), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5), consts.FacesUnion3Straight},
		{getPokers(3, 3, 4, 3, 3, 4, 4, 4, 5, 5, 5, 5), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 10), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 3, 3, 4, 4, 4, 4, 4), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 3, 4, 4, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9), consts.FacesUnion3Straight},
		{getPokers(3, 3, 3, 3, 3, 4, 4, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9), consts.FacesInvalid},
		{getPokers(3, 2, 2, 2, 3), consts.FacesUnion3},
		{getPokers(9, 9, 9, 8, 8, 8, 7, 7, 7, 11, 1, 15), consts.FacesUnion3Straight},
	}
	for _, testCase := range testCases {
		testParseFaces(t, testCase.pokers, testCase.actualType)
	}
}

func TestParseFaces2(t *testing.T) {
	pokers1 := getPokers(3, 3, 3, 3, 4, 4, 4, 4)
	faces1 := ParseFaces(pokers1, defaultRules)
	fmt.Println(faces1)

	//pokers1 := getPokers(8, 8, 8, 8, 2, 2, 2, 2)
	//faces1 := ParseFaces(pokers1, defaultRules)
	//pokers2 := getPokers(9, 9, 9, 9, 1, 2)
	//faces2 := ParseFaces(pokers2, defaultRules)
	//fmt.Println(faces1[0])
	//fmt.Println(faces2[0])
}

func TestParseFacesScore(t *testing.T) {
	testCases := []parseFacesCase{
		{pokers: getPokers(3, 3, 3, 3)},
		{pokers: getPokers(2, 2, 2, 2)},
		{pokers: getPokers(14, 14)},
		{pokers: getPokers(14, 15)},
		{pokers: getPokers(15, 15)},
		{pokers: getPokers(3, 3, 3, 3, 3)},
		{pokers: getPokers(2, 2, 2, 2, 2)},
		{pokers: getPokers(3, 3, 3, 3, 3, 3)},
		{pokers: getPokers(2, 2, 2, 2, 2, 2)},
		{pokers: getPokers(14, 14, 14)},
		{pokers: getPokers(14, 14, 15)},
		{pokers: getPokers(14, 15, 15)},
		{pokers: getPokers(15, 15, 15)},
		{pokers: getPokers(3, 3, 3, 3, 3, 3, 3)},
		{pokers: getPokers(2, 2, 2, 2, 2, 2, 2)},
		{pokers: getPokers(3, 3, 3, 3, 3, 3, 3, 3)},
		{pokers: getPokers(2, 2, 2, 2, 2, 2, 2, 2)},
		{pokers: getPokers(14, 14, 15, 15)},
	}
	preScore := int64(-1)
	for _, testCase := range testCases {
		list := ParseFaces(testCase.pokers, defaultRules)
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

func TestSort(t *testing.T) {
	pokers := getPokers(3, 8, 9)
	for i := range pokers {
		pokers[i].Val = defaultRules.Value(pokers[i].Key)
	}
	pokers.SetOaa(8, 9)
	pokers.SortByOaaValue()
	t.Log(pokers.String())

	pokers = getPokers(2, 5, 4, 9, 8)
	for i := range pokers {
		pokers[i].Val = defaultRules.Value(pokers[i].Key)
	}
	pokers.SetOaa(8, 9)
	pokers.SortByOaaValue()
	t.Log(pokers.String())
}
