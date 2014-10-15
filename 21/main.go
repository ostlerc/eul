package main

import (
	"fmt"
	"math"
)

var dc = make(map[int]int)

func factors(n int) []int {
	res := make([]int, 1, 20)
	res[0] = 1
	if n == 1 {
		return res
	}
	s := int(math.Sqrt(float64(n)))
	for i := 2; i < s; i++ {
		if n%i == 0 {
			res = append(res, i)
			res = append(res, n/i)
		}
	}

	if s*s == n {
		res = append(res, s)
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

func pop(n int) {
	for i := 0; i <= n; i++ {
		dc[i] = sumArray(factors(i))
	}
}

const c = 10000

func main() {
	pop(c)
	v := 0
	fmt.Println(220, dc[dc[220]])
	for i := 1; i < c; i++ {
		if x := dc[i]; x < c && i != x && i == dc[x] {
			fmt.Println(i)
			v += i
		}
	}

	fmt.Println(v, v/2)
}
