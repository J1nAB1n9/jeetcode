package example1669

import "fmt"

//	my test func
func AddValList(node *ListNode, arr []int) *ListNode {
	head := node
	for _, x := range arr {
		node.Next = AddVal(x)
		node = node.Next
	}
	return head.Next
}
func AddVal(x int) *ListNode {
	l := &ListNode{
		Val:  x,
		Next: nil,
	}
	return l
}

func ShowVal(l *ListNode, n bool) {
	fmt.Print(l.Val, " ")
	if l.Next != nil {
		ShowVal(l.Next, false)
	}

	if n {
		fmt.Print("\n")
	}
}

func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	return mergeInBetween(list1, a, b, list2)
}
