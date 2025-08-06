package example0904

func maxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func totalFruit(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	} else if len(fruits) == 1 {
		return 1
	}
	x, y, c := 0, 0, 0
	ans := 2
	for i := 1; i < len(fruits); i++ {
		if fruits[i] != fruits[x] && fruits[i] != fruits[y] {
			y = i
			x = c
			ans = maxx(ans, i-x+1)
		} else {
			ans = maxx(ans, i-x+1)
		}

		if fruits[i] != fruits[i-1] {
			// 最后处理
			c = i
		}
	}
	return ans
}
