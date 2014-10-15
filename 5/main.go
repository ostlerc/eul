package main

import (
	"flag"
	"fmt"
	"math"
)

func init() {
	flag.Parse()
}

func ed(n int) bool {
	for i := 1; i <= 20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	k := int64(math.Pow(float64(5050), float64(2)))
	sum := float64(1)
	for i := float64(2); i <= float64(100); i++ {
		sum += math.Pow(i, 2)
	}
	fmt.Println(sum, k, k-int64(sum))
}
