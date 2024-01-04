package example2397

import (
	"fmt"
	"math/bits"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumRows(matrix [][]int, numSelect int) int {
	m := len(matrix)
	n := len(matrix[0])
	zeroNum := 0
	var rows []int
	for i := 0; i < m; i++ {
		num := 0
		row := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				num++
				row |= 1 << (n - 1 - j)
			}
		}

		if num == 0 {
			zeroNum++
		} else if num <= numSelect {
			rows = append(rows, row)
		}
	}
	res, limit := 0, 1<<n
	for cur := 1; cur < limit; cur++ {
		if bits.OnesCount(uint(cur)) != numSelect {
			continue
		}
		t := 0
		for j := 0; j < len(rows); j++ {
			if (rows[j] & cur) == rows[j] {
				t++
			}
		}
		res = getMax(res, t)
	}
	fmt.Println(rows)
	fmt.Printf("%d %d\n", res, zeroNum)
	return res + zeroNum
}
