package holdem

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type HandType uint8
type HandValue int

const HighCard = HandType(0x00)
const OnePair = HandType(0x01)
const TwoPair = HandType(0x02)
const Trips = HandType(0x03)
const Straight = HandType(0x04)
const Flush = HandType(0x05)
const FullHouse = HandType(0x06)
const Quads = HandType(0x07)
const StraightFlush = HandType(0x08)
const RoyalFlush = HandType(0x09)

type Hand struct {
	Cards    [2]Card
	Score    HandValue
	HighHand [5]Card
}

type Round struct {
	Hands []Hand
	Flop  [3]Card
	Turn  Card
	River Card
}

func (round *Round) String() string {
	var buffer []string = make([]string, len(round.Hands)+3)
	buffer[0] = fmt.Sprintf("Draw:  %s %s %s", round.Flop[0], round.Flop[1], round.Flop[2])
	buffer[1] = fmt.Sprintf("Turn:  %s", round.Turn)
	buffer[2] = fmt.Sprintf("River: %s", round.River)
	for i, hand := range round.Hands {
		buffer[i+3] = fmt.Sprintf("Player %2d: %3s %3s (%06x using %3s %3s %3s %3s %3s)", i+1,
			hand.Cards[0], hand.Cards[1], hand.Score, hand.HighHand[0], hand.HighHand[1], hand.HighHand[2], hand.HighHand[3], hand.HighHand[4])
	}
	return strings.Join(buffer, "\n")
}

func NewRound(playerCount int8) (*Round, error) {
	if playerCount > 22 {
		return nil, fmt.Errorf("Too many players in a round: %d", playerCount)
	}
	if playerCount < 2 {
		return nil, fmt.Errorf("Too few players in a round: %d", playerCount)
	}
	shuffle := rand.Perm(len(Deck))
	round := &Round{
		Hands: make([]Hand, playerCount),
		Flop:  [3]Card{Deck[shuffle[0]], Deck[shuffle[1]], Deck[shuffle[2]]},
		Turn:  Deck[shuffle[3]],
		River: Deck[shuffle[4]],
	}
	for i := int8(0); i < playerCount; i++ {
		hand := [2]Card{Deck[shuffle[5+(i*2)]], Deck[shuffle[5+(i*2)+1]]}
		round.Hands[i].Cards = hand
		round.Hands[i].Score, round.Hands[i].HighHand = HighestHandValue([7]Card{
			hand[0], hand[1],
			round.Flop[0], round.Flop[1], round.Flop[2],
			round.Turn, round.River,
		})
	}
	return round, nil
}

func handValue(handType HandType, sortedValues [5]int) HandValue {
	return HandValue(0) |
		HandValue(handType)<<20 |
		HandValue(sortedValues[4])<<16 |
		HandValue(sortedValues[3])<<12 |
		HandValue(sortedValues[2])<<8 |
		HandValue(sortedValues[1])<<4 |
		HandValue(sortedValues[0])
}

func CalculateHandValue(hand [5]Card) HandValue {
	suit1 := hand[0] & 0x0F
	suit2 := hand[1] & 0x0F
	suit3 := hand[2] & 0x0F
	suit4 := hand[3] & 0x0F
	suit5 := hand[4] & 0x0F
	flush := bool(suit1 == suit2 && suit1 == suit3 && suit1 == suit4 && suit1 == suit5)

	cardValues := [5]int{
		int(hand[0] >> 4),
		int(hand[1] >> 4),
		int(hand[2] >> 4),
		int(hand[3] >> 4),
		int(hand[4] >> 4),
	}

	cardValuesSlice := sort.IntSlice(cardValues[:])
	cardValuesSlice.Sort()

	if cardValues[1] == cardValues[0]+1 && cardValues[2] == cardValues[1]+1 && cardValues[3] == cardValues[2]+1 {
		straight := false
		if cardValues[4] == cardValues[3]+1 {
			straight = true
		} else if cardValues[4] == 0x0E && cardValues[0] == 0x02 { //Low wheel
			cardValues = [5]int{0x01, 0x02, 0x03, 0x04, 0x05}
			straight = true
		}
		if straight && flush {
			if cardValues[4] == 0x0E {
				return handValue(RoyalFlush, cardValues)
			}
			return handValue(StraightFlush, cardValues)
		} else if straight {
			return handValue(Straight, cardValues)
		}
	}

	if flush {
		return handValue(Flush, cardValues)
	}

	// Check for Quads
	if cardValues[1] == cardValues[2] && cardValues[1] == cardValues[3] {

		if cardValues[0] == cardValues[1] {
			return handValue(Quads, cardValues)
			cardValuesSlice.Swap(0, 4)
		} else if cardValues[3] == cardValues[4] {
			return handValue(Quads, cardValues)
		} else {
			cardValuesSlice.Swap(1, 4)
			return handValue(Trips, cardValues)
		}

	}

	// Check for Full House
	if cardValues[0] == cardValues[1] && cardValues[3] == cardValues[4] {
		if cardValues[1] == cardValues[2] {
			cardValuesSlice.Swap(0, 4)
			cardValuesSlice.Swap(1, 3)
			return handValue(FullHouse, cardValues)
		} else if cardValues[2] == cardValues[3] {
			return handValue(FullHouse, cardValues)
		} else {
			// Swaping the middle card for the last card. Moves the two pairs in front for ranking.
			cardValuesSlice.Swap(2, 0)
			return handValue(TwoPair, cardValues)
		}
	}

	// Check for Trips (middle trip already covered in Quad check)
	if cardValues[0] == cardValues[1] && cardValues[1] == cardValues[2] {
		cardValuesSlice.Swap(0, 3)
		cardValuesSlice.Swap(1, 4)
		return handValue(Trips, cardValues)
	} else if cardValues[2] == cardValues[3] && cardValues[3] == cardValues[4] {
		return handValue(Trips, cardValues)
	}

	// Check for Two Pair (0+1 + 3+4 already covered by full house check)
	if cardValues[0] == cardValues[1] && cardValues[2] == cardValues[3] {
		cardValuesSlice.Swap(2, 4)
		cardValuesSlice.Swap(2, 0)
		return handValue(TwoPair, cardValues)
	} else if cardValues[1] == cardValues[2] && cardValues[3] == cardValues[4] {
		return handValue(TwoPair, cardValues)
	}

	// Check for Pair
	if cardValues[0] == cardValues[1] {
		cardValuesSlice.Swap(2, 4)
		cardValuesSlice.Swap(0, 4)
		cardValuesSlice.Swap(1, 3)
		return handValue(OnePair, cardValues)
	} else if cardValues[1] == cardValues[2] {
		cardValuesSlice.Swap(2, 4)
		cardValuesSlice.Swap(1, 3)
		return handValue(OnePair, cardValues)
	} else if cardValues[2] == cardValues[3] {
		cardValuesSlice.Swap(2, 4)
		return handValue(OnePair, cardValues)
	} else if cardValues[3] == cardValues[4] {
		return handValue(OnePair, cardValues)
	}

	return handValue(HighCard, cardValues)
}

var HandCombinations = [21][5]int8{
	[5]int8{1, 2, 3, 4, 5},
	[5]int8{1, 2, 3, 4, 6},
	[5]int8{1, 2, 3, 4, 7},
	[5]int8{1, 2, 3, 5, 6},
	[5]int8{1, 2, 3, 5, 7},
	[5]int8{1, 2, 3, 6, 7},
	[5]int8{1, 2, 4, 5, 6},
	[5]int8{1, 2, 4, 5, 7},
	[5]int8{1, 2, 4, 6, 7},
	[5]int8{1, 2, 5, 6, 7},
	[5]int8{1, 3, 4, 5, 6},
	[5]int8{1, 3, 4, 5, 7},
	[5]int8{1, 3, 4, 6, 7},
	[5]int8{1, 3, 5, 6, 7},
	[5]int8{1, 4, 5, 6, 7},
	[5]int8{2, 3, 4, 5, 6},
	[5]int8{2, 3, 4, 5, 7},
	[5]int8{2, 3, 4, 6, 7},
	[5]int8{2, 3, 5, 6, 7},
	[5]int8{2, 4, 5, 6, 7},
	[5]int8{3, 4, 5, 6, 7},
}

func HighestHandValue(cards [7]Card) (HandValue, [5]Card) {
	highest := HandValue(0)
	var highHand [5]Card
	for _, combination := range HandCombinations {
		hand := [5]Card{
			cards[combination[0]-1],
			cards[combination[1]-1],
			cards[combination[2]-1],
			cards[combination[3]-1],
			cards[combination[4]-1],
		}
		value := CalculateHandValue(hand)
		if value > highest {
			highest = value
			highHand = hand
		}
	}
	return highest, highHand
}
