package main

import "fmt"

var memo = make(map[int]map[int]int)
var cachedused = 0

func get(x, y int) (int, bool) {
	if _, ok := memo[x]; ok {
		if v, ok := memo[x][y]; ok {
			return v, true
		}
	}
	return 0, false
}

func insert(x, y, r int) {
	if _, ok := memo[x]; !ok {
		memo[x] = make(map[int]int)
	}
	if _, ok := memo[x]; !ok {
		memo[x] = make(map[int]int)
	}
	memo[x][y] = r
}

func p(x, y int) int {
	if x == 0 && y == 0 {
		return 0
	} else if x == 0 || y == 0 {
		return 1
	}
	if v, ok := get(x, y); ok {
		cachedused++
		return v
	}
	r := 0
	if x > 0 {
		r += p(x-1, y)
	}
	if y > 0 {
		r += p(x, y-1)
	}

	insert(x, y, r)

	return r
}

func main() {
	for i := 0; i < 21; i++ {
		cachedused = 0
		fmt.Println(i, p(i, i), cachedused)
	}
}
