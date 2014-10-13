package main

import "fmt"

var (
	//zero,one,two,three,four,five,six,seven,eight,nine
	//ones = [...]int{4, 3, 3, 5, 4, 4, 3, 5, 5, 4}
	ones = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	//teens = [...]int{3, 6, 6, 8, 8, 7, 7, 9, 9, 8}
	teens = [...]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	//tens = [...]int{-99999, -99999, 6, 6, 5, 5, 5, 7, 6, 6}
	tens = [...]string{"x", "x", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	//order = [...]int{7, 8}
	order = [...]string{"hundred", "thousand"}

	//and   = 3
	and = "and"
)

func str(n int) string {
	if n == 1000 {
		return ones[1] + order[1]
	}
	v := ""
	if n >= 100 {
		h := n / 100
		v += ones[h] + order[0]
		n -= h * 100
		if n > 0 {
			v += and
		}
	}

	if n > 19 { //tens
		t := n / 10
		v += tens[t]
		n -= t * 10
	} else if n >= 10 { //10-19
		v += teens[n-10]
		n = 0
	}
	if n > 0 {
		v += ones[n]
	}
	return v
}

func main() {
	total := 0
	for i := 1; i <= 1000; i++ {
		v := str(i)
		total += len(v)
		fmt.Println(i, v, total)
	}
	fmt.Println(total)
}
