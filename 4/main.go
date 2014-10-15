package main

import (
	"flag"
	"fmt"

	"strconv"
)

var (
	factor = flag.Int64("f", 13195, "factor value")
)

func init() {
	flag.Parse()
}

func isPal(n int) bool {
	s := strconv.Itoa(n)
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	a, b := 0, 0
	for i := 0; i < 999; i++ {
		for j := 0; j < 999; j++ {
			m := i * j
			if isPal(m) && m > max {
				max = m
				a, b = i, j
			}
		}
	}

	fmt.Println(max, a, b)
}
