package example0117

func findNext(pre *Node) *Node {
	if pre == nil {
		return nil
	}

	if pre.Left == nil && pre.Right == nil {
		return findNext(pre.Next)
	}

	if pre.Left != nil {
		return pre.Left
	}

	return pre.Right
}

func logic(root *Node) {
	if root == nil {
		return
	}

	ln := findNext(root.Next)

	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			root.Left.Next = ln
		}
	}

	if root.Right != nil {
		root.Right.Next = ln
	}

	logic(root.Right)
	logic(root.Left)
}

func connect(root *Node) *Node {
	logic(root)
	return root
}

//输入
//[2,1,3,0,7,9,1,2,null,1,0,null,null,8,8,null,null,null,null,7]
//输出
//[2,#,1,3,#,0,7,9,1,#,2,1,0,#,7,#]
//预期结果
//[2,#,1,3,#,0,7,9,1,#,2,1,0,8,8,#,7,#]
