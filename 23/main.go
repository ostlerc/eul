package main

import (
	"fmt"
	"math"
)

const c = 28123

var abundant = pop()

func factors(n int) []int {
	res := make([]int, 1)
	res[0] = 1
	s := int(math.Sqrt(float64(n)))
	for i := 2; i <= s; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i != n/i {
				res = append(res, n/i)
			}
		}
	}

	return res
}

func sumArray(a []int) int {
	v := 0
	for _, c := range a {
		v += c
	}
	return v
}

func pop() map[int]bool {
	a := make(map[int]bool, c)
	for i := 0; i < c; i++ {
		a[i] = sumArray(factors(i)) > i
	}
	return a
}

func summable(n int) bool {
	for i := 1; i <= n/2; i++ {
		if abundant[i] && abundant[n-i] {
			return true
		}
	}
	return false
}

func sum() int {
	res := 0
	for i := 1; i <= c; i++ {
		if !summable(i) {
			res += i
		}
	}
	return res
}

func main() {
	fmt.Println(sum())
}
