package math

import "math"

func Pow(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func Sum(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
