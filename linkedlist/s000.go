package linkedlist

//ListNode defines a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//CreateList create a list from an int slice
func CreateList(nums []int) *ListNode {
	head := &ListNode{}
	pre := head
	for _, num := range nums {
		node := ListNode{num, nil}
		pre.Next = &node
		pre = &node
	}
	return head.Next
}

//TraverseList transfer a list into an int slice
func TraverseList(head *ListNode) []int {
	nums := []int{}
	for ; head != nil; head = head.Next {
		nums = append(nums, head.Val)
	}
	return nums
}
