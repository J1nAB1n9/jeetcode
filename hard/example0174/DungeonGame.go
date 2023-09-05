package example0174

import (
	"math"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	n := len(dungeon)
	m := len(dungeon[0])

	var pd [][]int
	for i := 0; i < n; i++ {
		pd = append(pd, make([]int, m))
		for j := 0; j < m; j++ {
			pd[i][j] = math.MaxInt32
		}
	}

	pd[n-1][m-1] = getMax(1-dungeon[n-1][m-1], 1)

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i+1 >= n && j+1 >= m {
				continue
			}

			if i+1 >= n {
				pd[i][j] = getMax(pd[i][j+1]-dungeon[i][j], 1)
			} else if j+1 >= m {
				pd[i][j] = getMax(pd[i+1][j]-dungeon[i][j], 1)
			} else {
				pd[i][j] = getMax(getMin(pd[i+1][j], pd[i][j+1])-dungeon[i][j], 1)
			}
		}
	}

	//for i, v := range pd {
	//	for j, _ := range v {
	//		fmt.Printf("[%5d] ", pd[i][j])
	//	}
	//	fmt.Println()
	//}

	return pd[0][0]
}
