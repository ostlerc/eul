package main

import (
	"fmt"
	"math"
)

func factors(n int) int {
	count := 2
	s := int(math.Sqrt(float64(n)))
	for i := 2; i <= s; i++ {
		if n%i == 0 {
			count += 2
		}
	}

	if s*s == n {
		count--
	}
	return count
}

func main() {
	i := 1
	number := 0
	for factors(number) < 500 {
		number += i
		i++
	}
	fmt.Println(i, number)
}
