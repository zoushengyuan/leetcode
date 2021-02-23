package main

import (
	"fmt"
	. "leetcode/linkedlist"
)

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur *ListNode = head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func main() {
	head := CreateList([]int{1, 2, 3, 4, 5})
	head = reverseList(head)
	fmt.Println(TraverseList(head))
}
