package example2319

func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	for i, arr := range grid {
		for j, val := range arr {
			if i == j || n-i-1 == j {
				if val == 0 {
					return false
				}
			} else {
				if val != 0 {
					return false
				}
			}
		}
	}

	return true
}
