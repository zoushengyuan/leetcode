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
		//fmt.Println(lo, mid, hi)
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
	r := search(nums, target)
	if r == -1 {
		return []int{-1, -1}
	}
	start := r

	rs := r
	for rs > 1 {
		//fmt.Println(nums[:rs], target)

		rs = search(nums[:rs], target)
		//fmt.Println(rs)
		if rs == -1 {
			break
		}
		start = rs
	}

	end := r
	re := r
	n := len(nums) - 1
	for re < n {
		temp := search(nums[re+1:], target)
		//fmt.Println(re, nums[re+1:], temp)
		if temp == -1 {
			break
		}
		end = re + 1 + temp
		re = end
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

}
