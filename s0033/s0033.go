package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	n := len(nums)
	lo := 0
	hi := n - 1
	for lo <= hi {
		//fmt.Println(lo, hi)
		mid := lo + (hi-lo)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] <= nums[hi] {
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			if target < nums[mid] && target >= nums[lo] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}
	return -1
}

func main() {
	var nums []int
	var target int

	nums = []int{4, 5, 6, 7, 0, 1, 2, 3}
	target = 0
	fmt.Println(search(nums, target))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 3
	fmt.Println(search(nums, target))

	nums = []int{1}
	target = 0
	fmt.Println(search(nums, target))

}
