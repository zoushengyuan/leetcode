package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	n := len(s)
	isPalindrome := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
	}
	var longest string
	for size := 0; size < n; size++ {
		for lo := 0; lo < n; lo++ {
			hi := lo + size
			if hi >= n {
				continue
			}
			if size == 0 {
				isPalindrome[lo][hi] = true
			} else if size == 1 {
				isPalindrome[lo][hi] = s[lo] == s[hi]
			} else {
				isPalindrome[lo][hi] = isPalindrome[lo+1][hi-1] && s[lo] == s[hi]
			}
			if isPalindrome[lo][hi] && size+1 > len(longest) {
				longest = s[lo : hi+1]
			}
		}
	}
	return longest
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
}
