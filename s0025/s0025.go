package main

import (
	"fmt"
	. "leetcode/linkedlist"
)

func reverse(head *ListNode, k int) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	current := head
	//pre := &ListNode{0, head}
	var pre *ListNode = nil
	i := 0
	for i < k && current != nil {
		temp := current.Next
		current.Next = pre
		pre = current
		current = temp
		i++
	}
	head.Next = current
	tail := head
	head = pre
	if i < k {
		return reverse(head, i)
	}
	return head, tail
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	head, tail := reverse(head, k)
	for tail.Next != nil {
		tail.Next, tail = reverse(tail.Next, k)
	}
	return head
}

func main() {
	var head *ListNode

	head = CreateList([]int{1, 2, 3})
	head, _ = reverse(head, 3)
	fmt.Println(TraverseList(head))

	head = CreateList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	head = reverseKGroup(head, 3)
	fmt.Println(TraverseList(head))

	head = CreateList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	head = reverseKGroup(head, 3)
	fmt.Println(TraverseList(head))

	head = CreateList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	head = reverseKGroup(head, 3)
	fmt.Println(TraverseList(head))
}
