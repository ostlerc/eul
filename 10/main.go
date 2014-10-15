package main

import (
	"fmt"
	"math"
	"math/big"
)

func isPrime(n *big.Int) bool {
	sn := int64(math.Sqrt(float64(n.Int64())))
	t := big.NewInt(0)
	for i := int64(2); i <= sn; i++ {
		t.Mod(n, big.NewInt(i))
		if t.Int64() == 0 {
			return false
		}
	}

	return true
}

func sum(n int64) int64 {
	one := big.NewInt(1)
	s := big.NewInt(5)
	for i := big.NewInt(4); i.Int64() <= n; i.Add(i, one) {
		if isPrime(i) {
			s.Add(s, i)
		}
	}
	return s.Int64()
}

func main() {
	fmt.Println(sum(2000000))
}
