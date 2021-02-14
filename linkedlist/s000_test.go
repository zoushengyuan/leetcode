package linkedlist

import (
	"fmt"
	"testing"
)

func Test_createList(t *testing.T) {
	var nums []int
	var head *ListNode

	nums = []int{1, 2, 3, 4, 5}
	head = CreateList(nums)
	fmt.Println(TraverseList(head))
}

func Test_traverseList(t *testing.T) {
	var nums []int
	var head *ListNode

	nums = []int{}
	head = CreateList(nums)
	fmt.Println(TraverseList(head))
}
