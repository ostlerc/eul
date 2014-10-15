package main

import (
	"fmt"
	"math"
)

func trip(a, b, c int) bool {
	d, e, f := float64(a), float64(b), float64(c)
	return math.Pow(d, 2)+math.Pow(e, 2) == math.Pow(f, 2)
}

func main() {
	for i := 1; i < 998; i++ {
		for j := 1; j < 998; j++ {
			for k := 1; k < 998; k++ {
				if i+j+k == 1000 && trip(i, j, k) {
					fmt.Println(i, j, k, i*j*k)
				}
			}
		}
	}
	fmt.Println("done")
}
