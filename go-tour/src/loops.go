package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Sqrt(x float64) float64 {
	var z float64 = rand.Float64() + float64(1)
	z_next := z + float64(1)
	
	
	for diff := z_next - z; math.Floor(diff) != 0; {
		z_next = z - ((z*z - x) / (2 * z))
		diff = z_next - z
	}
	return z_next
}

func main() {
	fmt.Println(Sqrt(2))
}
