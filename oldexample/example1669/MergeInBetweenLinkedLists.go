package example1669

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
* 这题执行用时和内存消耗没有意义，量级都太小了
 */

//////////////////////////////////////////////////////
//	执行用时：84 ms, 在所有 Go 提交中击败了81.18% 的用户	//
//	内存消耗：7.3 MB, 在所有 Go 提交中击败了15.29% 的用户	//
//////////////////////////////////////////////////////
//func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
//	head := list1
//	i := 0
//	var nodeA *ListNode
//	var nodeB *ListNode
//
//	for ; list1 != nil; list1 = list1.Next {
//		if i == a-1 {
//			nodeA = list1
//		} else if i == b+1 {
//			nodeB = list1
//			break
//		}
//		i++
//	}
//
//	nodeA.Next = list2
//	for ; list2.Next != nil; list2 = list2.Next {
//		// Pass
//	}
//	list2.Next = nodeB
//	return head
//}

//////////////////////////////////////////////////////
//	执行用时：116 ms, 在所有 Go 提交中击败了5.88% 的用户	//
//	内存消耗：7.3 MB, 在所有 Go 提交中击败了43.53% 的用户	//
//////////////////////////////////////////////////////
//func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
//	head := list1
//	i := 0
//
//	for list1 != nil {
//		if i == a-1 {
//			tmpNode := list1.Next
//			list1.Next = list2
//			list1 = tmpNode
//			for ; list2.Next != nil; list2 = list2.Next {
//				// Pass
//			}
//			i++
//			continue
//		} else if i == b+1 {
//			list2.Next = list1
//			break
//		}
//		i++
//		list1 = list1.Next
//	}
//
//	return head
//}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	nodeA, nodeB := list1, list1
	for i := 1; i <= b; i++ {
		if i < a {
			nodeA = nodeA.Next
		}
		nodeB = nodeB.Next
	}

	nodeA.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = nodeB.Next

	return list1
}
