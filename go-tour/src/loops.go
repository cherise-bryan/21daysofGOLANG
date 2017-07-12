package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Sqrt(x float64) float64 {
	var z float64 = rand.Float64() + float64(1)
	var diff float64 = x

	for math.Floor(diff) != 0 {
		diff = (z - ((z*z - x) / 2 * z))
	}
	return diff
}

func main() {
	fmt.Println(Sqrt(2))
}
