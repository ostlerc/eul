package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	ctri [][]int
	memo = make(map[int]map[int]int)
)

func init() {
	dat, err := ioutil.ReadFile("triangle.txt")
	if err != nil {
		panic(err)
	}
	ctri = loadTri(string(dat))
}

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
		if len(a) > 0 {
			res = append(res, a)
		}
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
	/*	for _, l := range ctri {
		fmt.Println(l)
	}*/
	fmt.Println(path(0, 0))
}
