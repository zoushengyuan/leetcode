package main

import (
	"fmt"
	"sort"
)

func dif(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	r := 3001 //any one of nums <= 1000
	for i := 0; i < n-2; i++ {
		low := i + 1
		high := n - 1
		for high > low {
			s := nums[i] + nums[low] + nums[high]
			fmt.Print(nums[i], nums[low], nums[high], s)
			if s > target {
				if s-target < dif(r, target) {
					r = s
				}
				high--
			} else if s < target {
				if target-s < dif(r, target) {
					r = s
				}
				low++
			} else {
				return target
			}
			fmt.Println(r)
		}
	}
	return r
}

func main() {
	var nums []int
	var target int

	nums = []int{-1, 2, 1, -4}
	target = 1
	fmt.Println(threeSumClosest(nums, target))
}
