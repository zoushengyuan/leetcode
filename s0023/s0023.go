package main

import (
	"fmt"
)

/*
func findMin(lists []*ListNode) int {
	n := len(lists)
	if n == 0 {
		return -1
	}
	minVal := math.MaxInt64
	minPos := -1
	//var list *ListNode
	for i, list := range lists {
		if list != nil && list.Val < minVal {
			minVal = list.Val
			minPos = i
		}
	}
	return minPos
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for {
		pos := findMin(lists)
		if pos == -1 {
			p.Next = nil
			break
		}
		p.Next = lists[pos]
		p = p.Next
		lists[pos] = lists[pos].Next
		//fmt.Println(p.Val)
	}
	return head.Next
}
*/


func main() {
	list1 := createList([]int{1, 4, 5})
	list2 := createList([]int{1, 3, 4})
	list3 := createList([]int{2, 6})
	lists := []*ListNode{list1, list2, list3}
	//fmt.Println(findMin(lists))
	fmt.Println(traverseList(mergeKLists(lists)))
}
