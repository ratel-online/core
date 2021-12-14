package model

import (
	"bytes"
	"fmt"
	"github.com/ratel-online/core/consts"
	"github.com/ratel-online/core/util/arrays"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Pokers []Poker

type Poker struct {
	Key  int    `json:"key"`
	Val  int    `json:"val"`
	Type int    `json:"type"`
	Desc string `json:"desc"`
	Oaa  bool   `json:"oaa"`
}

type Faces struct {
	Keys   []int            `json:"keys"`
	Values []int            `json:"values"`
	Score  int64            `json:"score"`
	Type   consts.FacesType `json:"type"`
	Main   int              `json:"main"`
	Extra  int              `json:"extra"`
	HasOaa bool             `json:"hasOaa"`
}

func (f *Faces) SetValues(values []int) *Faces {
	f.Values = values
	return f
}

func (f *Faces) SetKeys(keys []int) *Faces {
	f.Keys = keys
	return f
}

func (f *Faces) SetMain(main int) *Faces {
	f.Main = main
	return f
}

func (f *Faces) SetExtra(extra int) *Faces {
	f.Extra = extra
	return f
}

func (f *Faces) SetScore(score int64) *Faces {
	f.Score = score
	return f
}

func (f *Faces) SetType(t consts.FacesType) *Faces {
	f.Type = t
	return f
}

func (f *Faces) Hash() string {
	buf := bytes.Buffer{}
	for _, k := range f.Keys {
		buf.WriteString(strconv.Itoa(k))
	}
	return buf.String()
}

func (f Faces) Compare(lastFaces Faces) bool {
	if f.Type == consts.FacesBomb {
		return f.Score > lastFaces.Score
	}
	if f.Type != lastFaces.Type {
		return false
	}
	return f.Score > lastFaces.Score && f.Main == lastFaces.Main && f.Extra == lastFaces.Extra
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

func (pokers Pokers) Keys() []int {
	keys := make([]int, 0)
	for _, poker := range pokers {
		keys = append(keys, poker.Key)
	}
	return keys
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

func (pokers Pokers) SetOaa(oaa ...int) {
	for i := range pokers {
		if arrays.Contains(oaa, pokers[i].Key) {
			pokers[i].Oaa = true
		}
	}
}

func (pokers Pokers) SortByOaaValue() {
	sort.Slice(pokers, func(i, j int) bool {
		if pokers[i].Oaa == pokers[j].Oaa {
			return pokers[i].Val < pokers[j].Val
		} else if pokers[i].Oaa {
			return false
		} else {
			return true
		}
	})
}

func (pokers Pokers) String() string {
	return pokers.OaaString()
}

func (pokers Pokers) OaaString() string {
	buf := bytes.Buffer{}
	for i := len(pokers) - 1; i >= 0; i-- {
		poker := pokers[i]
		flag := ""
		if poker.Oaa {
			flag = "*"
		}
		buf.WriteString(fmt.Sprintf("%s%v", flag, poker.Desc))
		if i != 0 {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}
