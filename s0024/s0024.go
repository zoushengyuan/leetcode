package main

import (
	"fmt"
)

func swapPairs(head *ListNode) *ListNode {
	p := head
	//var a, b *ListNode
	head = &ListNode{}
	pre := head
	head.Next = p
	for {
		if p == nil {
			break
		}
		a := p
		p = p.Next
		if p == nil {
			break
		}
		b := p
		p = p.Next
		pre.Next = b
		a.Next = b.Next
		b.Next = a
		pre = a
	}
	return head.Next
}

func main() {
	var head *ListNode

	head = createList([]int{1, 2, 3, 4})
	head = swapPairs(head)
	fmt.Println(traverseList(head))

	head = createList([]int{1})
	head = swapPairs(head)
	fmt.Println(traverseList(head))
}
