package example0023

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 原型方法mergeKLists00太慢了无法接受
//272ms
//击败 5.02%使用 Go 的用户
//
//4.98MB
//击败 66.07%使用 Go 的用户
func mergeKLists00(lists []*ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	pn := head

	size := len(lists)
	var maxx int
	var index int
aga:
	maxx = 0xffff
	index = -1
	for i := 0; i < size; i++ {
		if lists[i] == nil {
			continue
		}

		if maxx > lists[i].Val {
			maxx = lists[i].Val
			index = i
		}
	}

	if index != -1 {
		pn.Next = lists[index]
		pn = pn.Next
		lists[index] = lists[index].Next
		goto aga
	}
	return head.Next
}
