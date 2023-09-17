package example0337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dp(n *TreeNode) (int, int) {
	if n == nil {
		return 0, 0
	}

	ls, lgs := dp(n.Left)
	rs, rgs := dp(n.Right)

	// 两种状态，要嘛直接取子层的和，要嘛就取根节点+孙层的和
	sv := ls + rs
	gsv := lgs + rgs + n.Val
	if sv > gsv {
		return sv, sv
	} else {
		return gsv, sv
	}
}

// 基础Dp的变式，由于不太适合再用数组，所以直接用递归来写
func rob(root *TreeNode) int {
	ans, _ := dp(root)
	return ans
}
