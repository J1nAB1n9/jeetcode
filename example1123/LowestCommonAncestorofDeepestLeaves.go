package example1123

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsLeavesNode(t *TreeNode) bool {
	if t.Left == nil && t.Right == nil {
		return true
	}

	return false
}

var Hight int

func dg(t *TreeNode, treeHight int) (*TreeNode, int) {
	if IsLeavesNode(t) {
		return t, treeHight
	}

	if t.Left != nil && t.Right == nil {
		// 子节点肯定比父节点深
		return dg(t.Left, treeHight+1)
	} else if t.Left == nil && t.Right != nil {
		// 子节点肯定比父节点深
		return dg(t.Right, treeHight+1)
	} else if !IsLeavesNode(t) {
		lt, lh := dg(t.Left, treeHight+1)
		rt, rh := dg(t.Right, treeHight+1)

		// 返回更深的子节点
		if lh > rh {
			return lt, lh
		} else if rh > lh {
			return rt, rh
		}

		// 说t是平衡二叉树，所以返回t，树深返回子的树深表示自己的实力
		return t, lh
	}

	// 这里是为了防止编译报错，实际上是不可达的
	return t, treeHight

}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if IsLeavesNode(root) {
		return root
	}
	Hight = 0

	t, _ := dg(root, 1)
	return t
}
