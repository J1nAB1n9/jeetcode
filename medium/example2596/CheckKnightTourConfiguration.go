package example2596

type node struct {
	x int
	y int
}

func (n *node) IsRight(m node) bool {
	xv := n.x - m.x
	yv := n.y - m.y
	if xv*yv == 2 || xv*yv == -2 {
		return true
	}
	return false
}

func checkValidGrid(grid [][]int) bool {
	size := len(grid)
	nodes := make([]node, size*size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			index := grid[i][j]
			nodes[index].x = i
			nodes[index].y = j
		}
	}

	ju := node{
		x: 0,
		y: 0,
	}

	for i := 1; i < len(nodes); i++ {
		if !ju.IsRight(nodes[i]) {
			return false
		}

		ju = nodes[i]
	}

	return true
}
