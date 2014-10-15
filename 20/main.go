package main

import (
	"fmt"
	"math/big"
)

func sum(i *big.Int) uint64 {
	s := i.String()
	t := uint64(0)
	for _, c := range s {
		t += uint64(c - '0')
	}
	return t
}

func fact(n int) *big.Int {
	r := big.NewInt(int64(n))
	for i := 2; i < n; i++ {
		r.Mul(r, big.NewInt(int64(i)))
	}
	return r
}

func main() {
	fmt.Println(fact(100), sum(fact(100)))
}
