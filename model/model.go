package model

import (
	"math/rand"
	"time"
)

type AuthInfo struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
}

type Pokers []Poker

type Poker struct {
	Key  int    `json:"key"`
	Val  int    `json:"val"`
	Type int    `json:"type"`
	Desc string `json:"desc"`
}

type Faces struct {
	Types []int `json:"types"`
	Score int   `json:"score"`
}

func (pokers Pokers) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(pokers)
	r.Shuffle(l, func(i, j int) {
		pokers.Swap(i, j)
	})
}

func (pokers Pokers) Swap(i, j int) {
	pokers[i], pokers[j] = pokers[j], pokers[i]
}
