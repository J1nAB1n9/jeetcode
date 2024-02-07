package example2641

var sums []int

func CalcSum(root *TreeNode, height int) int {
	value := root.Val
	if height >= len(sums) {
		sums = append(sums, value)
	} else {
		sums[height] += value
	}

	left, right := 0, 0
	if root.Left != nil {
		left = CalcSum(root.Left, height+1)
	}

	if root.Right != nil {
		right = CalcSum(root.Right, height+1)
	}

	root.Val = left + right
	return value
}

func CalcAns(root *TreeNode, height int, fatherValue int) {
	ansValue := sums[height] - fatherValue

	if root.Left != nil {
		CalcAns(root.Left, height+1, root.Val)
	}

	if root.Right != nil {
		CalcAns(root.Right, height+1, root.Val)
	}

	root.Val = ansValue
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	sums = []int{}
	CalcSum(root, 0)
	CalcAns(root, 0, 0)

	root.Val = 0
	return root
}
