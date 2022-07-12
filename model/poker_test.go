package model

import (
	"fmt"
	"testing"
)

func TestPokers_Shuffle(t *testing.T) {
	pokers := make(Pokers, 0)
	for k := 1; k <= 54; k++ {
		pokers = append(pokers, Poker{
			Key: k,
			Val: 0,
		})
	}
	//pokers.Shuffle(len(pokers))
	for _, v := range pokers {
		fmt.Print(v.Key, " ")
	}
}
