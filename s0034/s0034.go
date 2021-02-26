package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	n := len(nums)
	lo := 0
	hi := n - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	start := search(nums, target)
	if start == -1 {
		return []int{-1, -1}
	}
	end := start
	r := start
	for r > 0 {
		r = search(nums[:r], target)
		if r == -1 {
			break
		}
		start = r
	}

	r = end
	n := len(nums) - 1
	for r < n {
		r = search(nums[r+1:], target)
		if r == -1 {
			break
		}
		end = end + 1 + r
		r = end
	}

	return []int{start, end}
}

func main() {
	var nums []int
	var target int

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 8
	fmt.Println(searchRange(nums, target))

	nums = []int{5, 7, 7, 8, 8, 8, 8, 10}
	target = 8
	fmt.Println(searchRange(nums, target))

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Println(searchRange(nums, target))

	nums = []int{}
	target = 0
	fmt.Println(searchRange(nums, target))

	nums = []int{1, 1, 2}
	target = 1
	fmt.Println(searchRange(nums, target))

}
