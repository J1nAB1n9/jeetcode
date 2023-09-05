package example0023

import (
	"fmt"
	"testing"
)

func Pnl(ln *ListNode) {
	for ln != nil {
		fmt.Printf("%d ", ln.Val)
		ln = ln.Next
	}
	fmt.Println()
}

func GetNode(list []int) *ListNode {
	head := &ListNode{}
	ph := head
	for _, v := range list {
		ph.Next = &ListNode{
			Val: v,
		}
		ph = ph.Next
	}

	return head.Next
}

func TestGetNode(t *testing.T) {
	Pnl(GetNode([]int{1, 4, 5}))
}

func TestExample(t *testing.T) {
	Pnl(mergeKLists([]*ListNode{GetNode([]int{1, 4, 5}), GetNode([]int{1, 3, 4}), GetNode([]int{2, 6})}))
	Pnl(mergeKLists([]*ListNode{GetNode([]int{}), GetNode([]int{})}))
	Pnl(mergeKLists([]*ListNode{GetNode([]int{})}))
}
