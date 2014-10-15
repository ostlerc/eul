package main

import "fmt"

func main() {
	a, b := int64(1), int64(2)
	sum := int64(0)
	for a < 4000000 && b < 4000000 {
		a, b = b, a+b
		if a%2 == 0 {
			sum += a
		}
	}
	fmt.Println(sum)
}
