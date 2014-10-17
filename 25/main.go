package main

import (
	"fmt"
	"math/big"
)

func fib() func() *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	return func() *big.Int {
		c := big.NewInt(0)
		c.Add(a, b)
		a, b = c, a
		return a
	}
}

func main() {
	f := fib()
	c := 0
	for {
		c++
		v := f()
		if len(v.String()) >= 1000 {
			fmt.Println(c, len(v.String()))
			break
		}
	}
}
