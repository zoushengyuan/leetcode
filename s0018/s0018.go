package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	r := [][]int{}
	prei := nums[0] - 1
	for i := 0; i < n-3; prei, i = nums[i], i+1 {
		if nums[i] == prei {
			continue
		}
		prej := nums[0] - 1
		for j := i + 1; j < n-2; prej, j = nums[j], j+1 {
			if nums[j] == prej {
				continue
			}
			low := j + 1
			high := n - 1
			for high > low {
				s := nums[i] + nums[j] + nums[low] + nums[high]
				if s > target {
					high--
				} else if s < target {
					low++
				} else {
					r = append(r, []int{nums[i], nums[j], nums[low], nums[high]})
					for low++; low < n && nums[low-1] == nums[low]; low++ {
						continue
					}
					high = n - 1
				}
			}
		}
	}
	return r
}

func main() {
	var nums []int
	var target int

	nums = []int{1, 0, -1, 0, -2, 2}
	target = 0
	fmt.Println(fourSum(nums, target))

	nums = []int{0, 0, 0, 0}
	target = 1
	fmt.Println(fourSum(nums, target))
}
