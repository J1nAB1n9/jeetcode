package example0107

type tree struct {
	maxHeight int
	queue     [][]int
}

func (t *tree) PushBack(val, height int) {
	if height > t.maxHeight {
		newList := make([][]int, height-t.maxHeight)
		t.queue = append(newList, t.queue...)
		t.maxHeight = height
	}

	index := t.maxHeight - height
	t.queue[index] = append(t.queue[index], val)
}

var tt tree

func dg(root *TreeNode, height int) {
	if root.Left != nil {
		dg(root.Left, height+1)
	}

	if root.Right != nil {
		dg(root.Right, height+1)
	}

	tt.PushBack(root.Val, height)
}

func levelOrderBottom(root *TreeNode) [][]int {
	tt.queue = [][]int{}
	tt.maxHeight = 0
	if root == nil {
		return tt.queue
	}

	dg(root, 1)

	return tt.queue
}
