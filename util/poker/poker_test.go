package poker

import (
	"fmt"
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var defaultRules = _defaultRules{}

type _defaultRules struct {
}

func (d _defaultRules) Score(key int) int {
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

func getPokers(keys ...int) model.Pokers {
	pokers := make(model.Pokers, 0)
	for _, k := range keys {
		pokers = append(pokers, model.Poker{
			Key:  k,
			Desc: desc(k),
			Type: 1,
		})
	}
	return pokers
}

func TestDistribute(t *testing.T) {
	pokersArr := Distribute(3, true)
	for _, pokers := range pokersArr {
		pokers.SortByKey()
		t.Log(pokers.String())
	}
	pokersArr = Distribute(5, true)
	for _, pokers := range pokersArr {
		pokers.SortByKey()
		t.Log(pokers.String())
	}
	pokersArr = Distribute(7, false)
	for _, pokers := range pokersArr {
		pokers.SortByKey()
		t.Log(pokers.String())
	}
}

type parseFacesCase struct {
	pokers     model.Pokers
	actualType consts.FacesType
}

func testParseFaces(t *testing.T, pokers model.Pokers, actualType consts.FacesType) {
	if !assert.Equal(t, ParseFaces(pokers, defaultRules).Type, actualType) {
		t.Log("err at", pokers.String())
	}
}

func TestParseFaces(t *testing.T) {
	testCases := []parseFacesCase{
		{getPokers(3), consts.FacesSingle},
		{getPokers(15), consts.FacesSingle},
		{getPokers(3, 3), consts.FacesDouble},
		{getPokers(3, 3, 3), consts.FacesTriple},
		{getPokers(3, 3, 3, 4), consts.FacesUnion},
		{getPokers(3, 3, 3, 4, 4), consts.FacesUnion},
		{getPokers(3, 3, 3, 4, 4, 5), consts.FacesInvalid},
		{getPokers(3, 3, 3, 3, 4), consts.FacesInvalid},
		{getPokers(3, 3, 3, 3, 4, 5), consts.FacesUnion},
		{getPokers(3, 3, 3, 3, 3, 4, 5), consts.FacesInvalid},
		{getPokers(3, 4, 5, 6, 7), consts.FacesStraight},
		{getPokers(3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1), consts.FacesStraight},
		{getPokers(2, 3, 4, 5, 6, 7), consts.FacesInvalid},
		{getPokers(10, 11, 12, 13, 1, 2), consts.FacesInvalid},
		{getPokers(3, 3, 4, 4, 5, 5), consts.FacesStraight},
		{getPokers(3, 3, 3, 4, 4, 4), consts.FacesStraight},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 4), consts.FacesUnionStraight},
		{getPokers(3, 3, 3, 3, 3, 4, 4, 4, 4, 4), consts.FacesUnionStraight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 6), consts.FacesUnionStraight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5, 6, 6), consts.FacesUnionStraight},
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
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5), consts.FacesUnionStraight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5, 5, 7, 7, 7), consts.FacesUnionStraight},
		{getPokers(3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 7, 7), consts.FacesUnionStraight},
		{getPokers(3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5), consts.FacesUnionStraight},
	}
	for _, testCase := range testCases {
		testParseFaces(t, testCase.pokers, testCase.actualType)
	}
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
		faces := ParseFaces(testCase.pokers, defaultRules)
		if faces.Score < preScore {
			t.Error(fmt.Sprintf("err score, pre %v should lt %v", preScore, faces.Score))
		}
		preScore = faces.Score
		t.Log(testCase.pokers.String(), faces.Score)
	}
}
