package main

import "fmt"

const (
	c   = 10
	end = 1000000
)

func perm(nums []byte) [][]byte {
	res := make([][]byte, 0)

	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return [][]byte{[]byte{nums[0], nums[1]}, []byte{nums[1], nums[0]}}
		} else {
			return [][]byte{[]byte{nums[1], nums[0]}, []byte{nums[0], nums[1]}}
		}
	}

	for i, v := range nums {
		next := make([]byte, len(nums))        //allocate new memory
		copy(next, nums)                       //copy nums
		next = append(next[:i], next[i+1:]...) //remove ith element
		for _, a := range perm(next) {
			t := append([]byte{v}, a...)
			res = append(res, t)
			if len(nums) == c && len(res) == end {
				fmt.Println("1 million!", res[end-1])
				return res
			}
		}
	}

	return res
}

func main() {
	initial := make([]byte, c)

	for i := 0; i < c; i++ {
		initial[i] = byte(i)
	}

	perm(initial)
}
