package darts

import "math"

func Score(x, y float64) int {
	uzunluk := math.Sqrt(x*x + y*y)

	if uzunluk <= 1 {
		return 10
	} else if uzunluk <= 5 {
		return 5
	} else if uzunluk <= 10 {
		return 1
	} else {
		return 0
	}
}
