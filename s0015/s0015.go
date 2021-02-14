package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	r := [][]int{}
	pre := nums[0] - 1
	for i := 0; i < n; pre, i = nums[i], i+1 {
		if nums[i] == pre {
			continue
		}
		low := i + 1
		high := n - 1
		for high > low {
			s := nums[i] + nums[low] + nums[high]
			if s > 0 {
				high--
			} else if s < 0 {
				low++
			} else {
				r = append(r, []int{nums[i], nums[low], nums[high]})
				for low++; low < n && nums[low-1] == nums[low]; low++ {
					continue
				}
				high = n - 1
			}
		}
	}
	return r
}

func main() {
	var nums []int
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
