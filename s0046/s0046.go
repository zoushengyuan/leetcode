package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}
	rs := [][]int{}
	for i := 0; i < n; i++ {
		nums[0], nums[i] = nums[i], nums[0]
		ps := permute(nums[1:])
		for _, p := range ps {
			r := make([]int, n)
			r[0] = nums[0]
			copy(r[1:], p)
			rs = append(rs, r)
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return rs
}

func main() {
	var nums []int

	nums = []int{1}
	fmt.Println(permute(nums))

	nums = []int{1, 2, 3}
	fmt.Println(permute(nums))
}
