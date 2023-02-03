package example1145

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Node struct {
	Val    int
	SonNun int
	Pre    *Node
	Left   *Node
	Right  *Node
}

var needNode *Node

func creatMyTree(thisNode *Node, srcTree *TreeNode, x int) int {
	if srcTree.Val == x {
		needNode = thisNode
	}

	thisNode.SonNun = 1
	if srcTree.Left != nil {
		thisNode.Left = &Node{
			Val:   srcTree.Left.Val,
			Pre:   thisNode,
			Left:  nil,
			Right: nil,
		}
		thisNode.SonNun += creatMyTree(thisNode.Left, srcTree.Left, x)
	}

	if srcTree.Right != nil {
		thisNode.Right = &Node{
			Val:   srcTree.Right.Val,
			Pre:   thisNode,
			Left:  nil,
			Right: nil,
		}
		thisNode.SonNun += creatMyTree(thisNode.Right, srcTree.Right, x)
	}
	return thisNode.SonNun
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	rootNode := &Node{
		Val:    root.Val,
		SonNun: 0,
	}
	rootNode.SonNun = creatMyTree(rootNode, root, x)
	// fmt.Printf("%d", rootNode.SonNun)
	if needNode.SonNun < (n/2 + 1) {
		return true
	}

	if needNode.Left.SonNun > (n / 2) {
		return true
	}

	if needNode.Right.SonNun > (n / 2) {
		return true
	}
	return false
}
