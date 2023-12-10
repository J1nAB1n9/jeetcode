package example1038

func nodeSun(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}

	root.Val += nodeSun(root.Right, val) + val
	return root.Val + nodeSun(root.Left, root.Val)
}

func bstToGst(root *TreeNode) *TreeNode {
	nodeSun(root, 0)
	return root
}
