package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	i := n - 2
	for ; i >= 0; i-- {
		//fmt.Println(i)
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i == -1 {
		sort.Ints(nums)
		return
	}
	//fmt.Println(i)

	for j := n - 1; j >= 0; j-- {
		if nums[j] > nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
			//sort.Sort(sort.Reverse(sort.IntSlice(nums[i+1:])))
			sort.Ints(nums[i+1:])
			return
		}
	}
}

func main() {
	nums := []int{3, 2, 1} // unsorted
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1} // unsorted
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 1, 5} // unsorted
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 2, 3} // unsorted
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 3, 2} // unsorted
	nextPermutation(nums)
	fmt.Println(nums)
}
