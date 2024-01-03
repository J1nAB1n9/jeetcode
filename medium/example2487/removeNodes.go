package example2487

import "fmt"

func removeNodes(head *ListNode) *ListNode {
	stack := []int{}
	for pp := head; pp != nil; pp = pp.Next {
		stackLen := len(stack)
		if stackLen == 0 {
			stack = append(stack, pp.Val)
		} else if stack[stackLen-1] >= pp.Val {
			stack = append(stack, pp.Val)
		} else {
			for len(stack) != 0 {
				if stack[len(stack)-1] >= pp.Val {
					stack = append(stack, pp.Val)
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}

			if len(stack) == 0 {
				stack = append(stack, pp.Val)
			}
		}
	}

	fmt.Println(stack)
	return head
}
