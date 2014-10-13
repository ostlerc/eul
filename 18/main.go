package main

import (
	"fmt"
	"strings"
)

const tri = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

/*const tri = `3
7 4
2 4 6
8 5 9 3`*/

var ctri = loadTri(tri)
var memo = make(map[int]map[int]int)

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

func loadTri(s string) [][]int {
	res := make([][]int, 0, 50)
	var t int
	for i, l := range strings.Split(s, "\n") {
		a := make([]int, 0, i+1)
		r := strings.NewReader(l)
		for {
			if _, err := fmt.Fscanf(r, "%d", &t); err != nil {
				break
			}
			a = append(a, t)
		}
		res = append(res, a)
	}

	return res
}

func path(n, i int) int {
	if i > n {
		return 0
	}
	if n == len(ctri)-1 {
		return ctri[n][i]
	}
	if v, ok := get(n, i); ok {
		return v
	}
	v := path(n+1, i)
	if v == 0 {
		panic("zero")
	}
	if t := path(n+1, i+1); v < t {
		v = t
	}
	v += ctri[n][i]
	insert(n, i, v)
	return v
}

func main() {
	for _, l := range ctri {
		fmt.Println(l)
	}
	fmt.Println(path(0, 0))
}
