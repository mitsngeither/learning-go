package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	for x := 1.0; x <= 5; x++ {
		fmt.Println(Sqrt(x), math.Sqrt(x))
	}
}
