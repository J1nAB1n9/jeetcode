package example2331

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func evaluateTree(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		if node.Val == 1 {
			return true
		} else {
			return false
		}
	}

	switch node.Val {
	case 2:
		return (evaluateTree(node.Left) || evaluateTree(node.Right))
	case 3:
		return (evaluateTree(node.Left) && evaluateTree(node.Right))
	default:
		return true
	}

	return true
}
