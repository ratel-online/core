package holdem

import "log"

type Card uint8

const AceSpades = Card(0xE1)
const TwoSpades = Card(0x21)
const ThreeSpades = Card(0x31)
const FourSpades = Card(0x41)
const FiveSpades = Card(0x51)
const SixSpades = Card(0x61)
const SevenSpades = Card(0x71)
const EightSpades = Card(0x81)
const NineSpades = Card(0x91)
const TenSpades = Card(0xA1)
const JackSpades = Card(0xB1)
const QueenSpades = Card(0xC1)
const KingSpades = Card(0xD1)

const AceClubs = Card(0xE2)
const TwoClubs = Card(0x22)
const ThreeClubs = Card(0x32)
const FourClubs = Card(0x42)
const FiveClubs = Card(0x52)
const SixClubs = Card(0x62)
const SevenClubs = Card(0x72)
const EightClubs = Card(0x82)
const NineClubs = Card(0x92)
const TenClubs = Card(0xA2)
const JackClubs = Card(0xB2)
const QueenClubs = Card(0xC2)
const KingClubs = Card(0xD2)

const AceHearts = Card(0xE3)
const TwoHearts = Card(0x23)
const ThreeHearts = Card(0x33)
const FourHearts = Card(0x43)
const FiveHearts = Card(0x53)
const SixHearts = Card(0x63)
const SevenHearts = Card(0x73)
const EightHearts = Card(0x83)
const NineHearts = Card(0x93)
const TenHearts = Card(0xA3)
const JackHearts = Card(0xB3)
const QueenHearts = Card(0xC3)
const KingHearts = Card(0xD3)

const AceDiamonds = Card(0xE4)
const TwoDiamonds = Card(0x24)
const ThreeDiamonds = Card(0x34)
const FourDiamonds = Card(0x44)
const FiveDiamonds = Card(0x54)
const SixDiamonds = Card(0x64)
const SevenDiamonds = Card(0x74)
const EightDiamonds = Card(0x84)
const NineDiamonds = Card(0x94)
const TenDiamonds = Card(0xA4)
const JackDiamonds = Card(0xB4)
const QueenDiamonds = Card(0xC4)
const KingDiamonds = Card(0xD4)

var Deck = [52]Card{
	AceSpades, TwoSpades, ThreeSpades, FourSpades, FiveSpades, SixSpades, SevenSpades, EightSpades, NineSpades, TenSpades, JackSpades, QueenSpades, KingSpades,
	AceClubs, TwoClubs, ThreeClubs, FourClubs, FiveClubs, SixClubs, SevenClubs, EightClubs, NineClubs, TenClubs, JackClubs, QueenClubs, KingClubs,
	AceHearts, TwoHearts, ThreeHearts, FourHearts, FiveHearts, SixHearts, SevenHearts, EightHearts, NineHearts, TenHearts, JackHearts, QueenHearts, KingHearts,
	AceDiamonds, TwoDiamonds, ThreeDiamonds, FourDiamonds, FiveDiamonds, SixDiamonds, SevenDiamonds, EightDiamonds, NineDiamonds, TenDiamonds, JackDiamonds, QueenDiamonds, KingDiamonds,
}

func (card Card) String() string {
	var cardNumber, cardSuite string
	num := card >> 4
	switch num {
	case 0x0E:
		cardNumber = "A"
	case 0x02:
		cardNumber = "2"
	case 0x03:
		cardNumber = "3"
	case 0x04:
		cardNumber = "4"
	case 0x05:
		cardNumber = "5"
	case 0x06:
		cardNumber = "6"
	case 0x07:
		cardNumber = "7"
	case 0x08:
		cardNumber = "8"
	case 0x09:
		cardNumber = "9"
	case 0x0A:
		cardNumber = "10"
	case 0x0B:
		cardNumber = "J"
	case 0x0C:
		cardNumber = "Q"
	case 0x0D:
		cardNumber = "K"
	default:
		log.Panicf("Invalid Card: %02x Number: %02x\n", uint8(card), cardNumber)
	}

	suit := card & 0x0F
	switch suit {
	case 0x01:
		cardSuite = "s"
	case 0x02:
		cardSuite = "c"
	case 0x03:
		cardSuite = "h"
	case 0x04:
		cardSuite = "d"
	default:
		log.Panicf("Invalid Card: %02x Suit: %02x\n", uint8(card), cardSuite)
	}

	return cardNumber + cardSuite
}
