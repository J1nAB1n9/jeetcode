package example2415

func reverseOddLevels(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	size := 0

	for size < len(queue) {
		node := queue[size]
		size++
		if node.Left == nil {
			continue
		}

		queue = append(queue, node.Left, node.Right)
	}

	time := 0
	leftIndex := 0
	rightIndex := 0
	for true {
		leftIndex = leftIndex*2 + 1
		rightIndex = rightIndex*2 + 2
		time++
		if time%2 == 0 {
			continue
		}

		if leftIndex >= size || rightIndex >= size {
			break
		}

		tmp := queue[leftIndex : rightIndex+1]
		for j, k := 0, len(tmp)-1; j < k; {
			tmp[j], tmp[k] = tmp[k], tmp[j]
			j++
			k--
		}
	}

	for i, v := range queue {
		//fmt.Printf("%d ", v.Val)
		if i*2+1 >= len(queue) {
			v.Left = nil
			v.Right = nil
			continue
		}

		queue[i].Left = queue[i*2+1]
		queue[i].Right = queue[i*2+2]
	}
	return root
}
