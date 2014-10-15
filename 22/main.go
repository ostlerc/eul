package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func score(s string) int64 {
	var res int64
	for _, c := range s {
		res += int64(c - 'A' + 1)
	}
	return res
}

func main() {
	dat, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		panic(err)
	}
	names := sort.StringSlice(strings.Split(strings.TrimSuffix(string(dat), "\n"), ","))
	sort.Sort(names)
	var scores int64
	for i, n := range names {
		s := int64(i+1) * score(n)
		scores += s
	}
	fmt.Println(scores)
}
