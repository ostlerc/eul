package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
)

var (
	factor = flag.Int64("f", 13195, "factor value")
)

func init() {
	flag.Parse()
}

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

func main() {
	sf := int64(math.Sqrt(float64(*factor)))
	fmt.Println(*factor, sf)
	bi := big.NewInt(*factor)
	set := make([]int64, 0, 50)
	t := big.NewInt(0)
	for i := int64(2); i < sf; i++ {
		bt := big.NewInt(int64(i))
		t.Mod(bi, bt)
		if t.Int64() == 0 && isPrime(bt) {
			set = append(set, i)
		}
	}
	fmt.Println(set)
}
