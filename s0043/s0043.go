package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	nums1 := []byte(num1)
	n1 := len(num1)
	charTodigit(nums1, n1)

	nums2 := []byte(num2)
	n2 := len(num2)
	charTodigit(nums2, n2)

	product := make([]byte, n1+n2)
	sum := make([]byte, n1+n2)

	for i := 0; i < n2; i++ {
		zeros(product)
		multiplyOneDigital(nums1, nums2[i], product[i:n1+i+1])
		add(sum, product, n1+n2)
	}

	return toString(sum, n1+n2)
}

func charTodigit(nums []byte, n int) {
	for i := 0; i < n; i++ {
		nums[i] -= '0'
	}
}

func toString(nums []byte, n int) string {
	i := 0
	for ; i < n-1; i++ {
		if nums[i] != 0 {
			break
		}
	}
	for j := i; j < n; j++ {
		nums[j] += '0'
	}

	return string(nums[i:])
}

func zeros(nums []byte) {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = 0
	}
}

func add(num1, num2 []byte, n int) {
	var carry byte = 0
	for i := n - 1; i >= 0; i-- {
		t := num1[i] + num2[i] + carry
		num1[i] = t % 10
		carry = t / 10
	}
}

//len(product) = len(nums) + 1
func multiplyOneDigital(nums []byte, d byte, product []byte) {
	var carry byte = 0
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		x := nums[i]*d + carry
		product[i+1] = x % 10
		carry = x / 10
	}
	product[0] = carry
}

func main() {
	var num1, num2 string

	num1 = "123"
	num2 = "0"
	fmt.Println(multiply(num1, num2))

	num1 = "12"
	num2 = "12"
	fmt.Println(multiply(num1, num2))

	num1 = "99"
	num2 = "2"
	fmt.Println(multiply(num1, num2))
}
