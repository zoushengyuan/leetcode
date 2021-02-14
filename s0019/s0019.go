package main

import (
	"fmt"
)

//ListNode defines a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre, cur, p *ListNode
	p = head
	for i := 1; p != nil && i < n; p = p.Next {
		i++
	}
	pre = nil
	cur = head
	//fmt.Println(p.Val)
	for p = p.Next; p != nil; p = p.Next {
		pre = cur
		cur = cur.Next
	}
	//fmt.Println(pre.Val, cur.Val)

	if pre != nil && cur != nil {
		pre.Next = cur.Next
		return head
	} else if pre == nil {
		return cur.Next
	}
	return nil
}

func main() {
	fmt.Println()
}
