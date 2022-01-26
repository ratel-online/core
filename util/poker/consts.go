package poker

import (
	"github.com/ratel-online/core/model"
	"github.com/ratel-online/core/util/strings"
	"math/rand"
	"time"
)

var (
	base      = make(model.Pokers, 0)
	desc      = map[int]string{}
	keysAlias = map[int][]string{}
	aliasKeys = map[string]int{}
)

func init() {
	for k := 1; k <= 15; k++ {
		switch k {
		case 1:
			desc[k] = "A"
			keysAlias[k] = []string{"1", "a", "A"}
		case 10:
			desc[k] = "10"
			keysAlias[k] = []string{"0", "t", "T"}
		case 11:
			desc[k] = "J"
			keysAlias[k] = []string{"j", "J"}
		case 12:
			desc[k] = "Q"
			keysAlias[k] = []string{"q", "Q"}
		case 13:
			desc[k] = "K"
			keysAlias[k] = []string{"k", "K"}
		case 14:
			desc[k] = "S"
			keysAlias[k] = []string{"s", "S"}
		case 15:
			desc[k] = "X"
			keysAlias[k] = []string{"x", "X"}
		default:
			desc[k] = strings.String(k)
			keysAlias[k] = []string{strings.String(k)}
		}
	}
	for k, aliases := range keysAlias {
		for _, alias := range aliases {
			aliasKeys[alias] = k
		}
	}
	for k := 1; k <= 13; k++ {
		for t := 1; t <= 4; t++ {
			base = append(base, model.Poker{Key: k, Val: 0, Desc: desc[k]})
		}
	}
	for k := 14; k <= 15; k++ {
		base = append(base, model.Poker{Key: k, Val: 0, Desc: desc[k]})
	}
}

func GetKey(alias string) int {
	return aliasKeys[alias]
}

func GetAlias(key int) string {
	if len(keysAlias[key]) > 0 {
		return keysAlias[key][0]
	}
	return ""
}

func GetDesc(key int) string {
	return desc[key]
}

func GetDontShuffleBase() model.Pokers {
	base := model.Pokers{}
	keys := make([]int, 0)
	for k := 1; k <= 15; k++ {
		keys = append(keys, k)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
	for _, k := range keys {
		if k <= 13 {
			for t := 1; t <= 4; t++ {
				base = append(base, model.Poker{Key: k, Val: 0, Desc: desc[k]})
			}
		} else {
			base = append(base, model.Poker{Key: k, Val: 0, Desc: desc[k]})
		}
	}
	return base
}
