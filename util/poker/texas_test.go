package poker

import (
	"github.com/ratel-online/core/model"
	"testing"
)

func TestParseTexasFaces(t *testing.T) {
	testCase := []struct {
		hand, board model.Pokers
		expected    *model.TexasFaces
	}{
		{
			hand:     model.Pokers{{Key: 1, Suit: model.Club}, {Key: 2, Suit: model.Club}},
			board:    model.Pokers{{Key: 3, Suit: model.Club}, {Key: 4, Suit: model.Club}, {Key: 5, Suit: model.Club}, {Key: 6, Suit: model.Club}, {Key: 7, Suit: model.Club}},
			expected: &model.TexasFaces{Type: model.TexasFacesTypeStraightFlush},
		},
		{
			hand:     model.Pokers{{Key: 1, Suit: model.Club}, {Key: 2, Suit: model.Club}},
			board:    model.Pokers{{Key: 3, Suit: model.Club}, {Key: 4, Suit: model.Diamond}, {Key: 5, Suit: model.Club}, {Key: 6, Suit: model.Club}, {Key: 8, Suit: model.Club}},
			expected: &model.TexasFaces{Type: model.TexasFacesTypeFlush},
		},
	}
	for _, tc := range testCase {
		actual, err := ParseTexasFaces(tc.hand, tc.board)
		if err != nil {
			t.Fatal(err)
		}
		if actual.Type != tc.expected.Type {
			t.Fatalf("expected %v, got %v", tc.expected.Type, actual.Type)
		}
	}
}
