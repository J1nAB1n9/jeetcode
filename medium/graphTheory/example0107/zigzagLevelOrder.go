package example0107

type tree struct {
	maxHeight int
	queue     [][]int
}

func (t *tree) PushBack(val, height int) {
	if height > t.maxHeight {
		newList := make([][]int, height-t.maxHeight)
		t.queue = append(t.queue, newList...)
		t.maxHeight = height
	}

	t.queue[height-1] = append(t.queue[height-1], val)
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	tt.maxHeight = 0
	tt.queue = [][]int{}
	if root == nil {
		return tt.queue
	}

	dg(root, 1)

	for i, arr := range tt.queue {
		if i%2 == 1 {
			for j := 0; j < len(arr)/2; j++ {
				r := len(arr) - 1 - j
				tt.queue[i][j], tt.queue[i][r] = tt.queue[i][r], tt.queue[i][j]
			}
		}
	}

	return tt.queue
}
