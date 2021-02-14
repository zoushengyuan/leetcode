package main

import (
	"fmt"
	. "leetcode/linkedlist"
)

func main() {
	var head *ListNode
	head = CreateList([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(TraverseList(head))
}
