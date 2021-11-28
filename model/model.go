package model

import (
	"bytes"
	"fmt"
	"github.com/ratel-online/core/consts"
	"math/rand"
	"sort"
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
	Counts []int            `json:"types"`
	Score  int64            `json:"score"`
	Type   consts.FacesType `json:"type"`
}

func (f *Faces) SetScore(score int64) *Faces {
	f.Score = score
	return f
}

func (f *Faces) SetType(t consts.FacesType) *Faces {
	f.Type = t
	return f
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

func (pokers Pokers) SortByKey() {
	sort.Slice(pokers, func(i, j int) bool {
		return pokers[i].Key < pokers[j].Key
	})
}

func (pokers Pokers) SortByValue() {
	sort.Slice(pokers, func(i, j int) bool {
		return pokers[i].Val < pokers[j].Val
	})
}

func (pokers Pokers) String() string {
	buf := bytes.Buffer{}
	for _, poker := range pokers {
		buf.WriteString(fmt.Sprintf("%v ", poker.Key))
	}
	return buf.String()
}
