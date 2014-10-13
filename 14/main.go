package main

import "fmt"

func chain(n int) int {
	count := 0
	for n != 1 {
		if n%2 != 0 { //odd
			n = 3*n + 1
		} else {
			n >>= 1
		}
		count++
	}
	return count
}

func main() {
	max := 0
	maxi := 0
	for i := 2; i < 1000000; i++ {
		v := chain(i)
		if v > max {
			max = v
			maxi = i
		}
		if i%10000 == 0 {
			fmt.Println(i / 10000)
		}
		//fmt.Println(i, v)
	}
	fmt.Println(max, maxi)
}
