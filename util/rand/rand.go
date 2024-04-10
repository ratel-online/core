package rand

import (
	"math/rand"
	"time"
)

func Intn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}
