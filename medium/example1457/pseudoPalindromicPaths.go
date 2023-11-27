package example1457

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ans int

func nextNode(root *TreeNode, mark int) {
	mark ^= (1 << root.Val)
	// fmt.Println(mark)
	if root.Left == nil && root.Right == nil {
		// 找完了
		lowBit := (mark & -mark)
		if lowBit == mark {
			ans++
		}
		return
	}

	if root.Left != nil {
		nextNode(root.Left, mark)
	}
	if root.Right != nil {
		nextNode(root.Right, mark)
	}
}

func pseudoPalindromicPaths(root *TreeNode) int {
	ans = 0
	nextNode(root, 0)
	return ans
}
