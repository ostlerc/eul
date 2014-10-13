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

func main() {
	bits := []big.Word{}
	i := 0
	for ; i < 1000-64; i += 64 {
		bits = append(bits, 0)
	}
	bits = append(bits, 1<<uint32(1000-i))
	n := big.NewInt(0).SetBits(bits)
	fmt.Println(i, bits, n, sum(n))
}
