package example2415

func reverseOddLevels(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	size := 0

	for size < len(queue) {
		node := queue[size]
		if node.Left == nil {
			continue
		}

		queue = append(queue, node.Left, node.Right)
		size++
	}

	for i := 0; (1 << i) < size; i += 2 {
		leftIndex := 1 << i
		rightIndex := (1 << (i + 1)) - 1
		if rightIndex >= size {
			rightIndex = size - 1
		}

		tmp := queue[leftIndex : rightIndex+1]
		for j, k := 0, len(tmp)-1; j < k; {
			tmp[j], tmp[k] = tmp[k], tmp[j]
			j++
			k--
		}
	}

	for i, v := range queue {
		if i*2+1 < len(queue) {
			v.Left = nil
			v.Right = nil
			continue
		}

		queue[i].Left = queue[i*2+1]
		queue[i].Right = queue[i*2+2]
	}
	return root
}
