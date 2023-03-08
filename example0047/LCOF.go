package example0047

func GetMax(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxValue(grid [][]int) int {
	nSize := len(grid)

	if nSize == 0 {
		return 0
	}
	mSize := len(grid[0])
	var pd [][]int
	for i := 0; i < nSize; i++ {
		pdd := make([]int, mSize)
		pd = append(pd, pdd)
	}

	pd[0][0] = grid[0][0]
	for i := 0; i < nSize; i++ {
		for j := 0; j < mSize; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				pd[i][j] = grid[i][j] + pd[i][j-1]
			} else if j == 0 {
				pd[i][j] = grid[i][j] + pd[i-1][j]
			} else {
				pd[i][j] = GetMax(pd[i-1][j]+grid[i][j], pd[i][j-1]+grid[i][j])
			}

		}
	}

	return pd[nSize-1][mSize-1]
}
