package main

import (
	"fmt"
	. "leetcode/linkedlist"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	current := &head
	for {
		if l1 == nil {
			current.Next = l2
			break
		} else if l2 == nil {
			current.Next = l1
			break
		} else if l1.Val > l2.Val {
			current.Next = l2
			current = l2
			l2 = l2.Next
		} else {
			current.Next = l1
			current = l1
			l1 = l1.Next
		}
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	l1 := lists[0]
	for _, l2 := range lists[1:] {
		l1 = mergeTwoLists(l1, l2)
	}
	return l1
}

func main() {
	list1 := CreateList([]int{1, 4, 5})
	list2 := CreateList([]int{1, 3, 4})
	list3 := CreateList([]int{2, 6})
	lists := []*ListNode{list1, list2, list3}
	fmt.Println(TraverseList(mergeKLists(lists)))
}
