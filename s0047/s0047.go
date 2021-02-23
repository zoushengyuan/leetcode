package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	//fmt.Println(nums)
	return permute(nums)
}

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}
	rs := [][]int{}
	pre := -11
	for i := 0; i < n; i++ {
		if pre == nums[i] {
			continue
		}
		others := make([]int, n-1)
		copy(others, nums[0:i])
		copy(others, nums[i+1:])
		fmt.Print(others)
		ps := permute(others[1:])
		for _, p := range ps {
			r := make([]int, n)
			r[0] = nums[0]
			copy(r[1:], p)
			rs = append(rs, r)
			fmt.Println(rs)
		}
		pre = nums[i]
	}
	return rs
}

func main() {
	var nums []int

	nums = []int{1, 2}
	ps := permuteUnique(nums)
	fmt.Println(ps)
	fmt.Println(len(ps))

	//nums = []int{2, 2, 2}
	//fmt.Println(permuteUnique(nums))
}
