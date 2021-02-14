package main

//ListNode defines a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	head := &ListNode{}
	pre := head
	for _, num := range nums {
		node := ListNode{num, nil}
		pre.Next = &node
		pre = &node
	}
	return head.Next
}

func traverseList(head *ListNode) []int {
	nums := []int{}
	for ; head != nil; head = head.Next {
		nums = append(nums, head.Val)
	}
	return nums
}
