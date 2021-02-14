package main

import (
	"fmt"
)

func reverse(lead *ListNode, head *ListNode, k int) *ListNode {
	p := head
	pre := lead
	i := 0
	for i < k && p != nil {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
		i++
	}
	lead.Next = pre
	head.Next = p
	return head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}

func main() {
	var head *ListNode

	lead := &ListNode{}

	head = createList([]int{1, 2, 3, 4})
	lead.Next = head
	reverse(lead, head, 2)
	fmt.Println(traverseList(lead.Next))
	//fmt.Println(traverseList(head))
}
