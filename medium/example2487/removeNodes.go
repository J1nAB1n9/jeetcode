package example2487

func removeNodes(head *ListNode) *ListNode {
	stack := []*ListNode{}
	for pp := head; pp != nil; pp = pp.Next {
		stackLen := len(stack)
		if stackLen == 0 {
			stack = append(stack, pp)
		} else if stack[stackLen-1].Val >= pp.Val {
			stack = append(stack, pp)
		} else {
			for len(stack) != 0 {
				if stack[len(stack)-1].Val >= pp.Val {
					stack = append(stack, pp)
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			}

			if len(stack) == 0 {
				stack = append(stack, pp)
			}
		}
	}

	node := &ListNode{}
	head = node
	for _, v := range stack {
		node.Next = v
		node = node.Next
	}
	return head.Next
}
