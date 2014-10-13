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

func pat(n int) int64 {
	found := 0
	one := big.NewInt(1)
	last := big.NewInt(0)
	for i := big.NewInt(2); found != n; i.Add(i, one) {
		if isPrime(i) {
			last.SetInt64(i.Int64())
			found++
		}
	}
	return last.Int64()
}

func main() {
	fmt.Println(pat(10001))
}
