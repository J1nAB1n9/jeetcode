package example0062

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func U(obstacleGrid [][]int) int {
	return uniquePathsWithObstacles(obstacleGrid)
}

// 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			continue
		}
		dp[i][0] = dp[i-1][0]
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			continue
		}
		dp[0][i] = dp[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		fmt.Printf("%d ", dp[i][j])
	//	}
	//	fmt.Printf("\n")
	//}

	return dp[m-1][n-1]
}

// 63. 不同路径 II
// 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// 输出：2
//
// 输入：obstacleGrid = [[0,1],[0,0]]
// 输出：1
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1

// example 120
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	dp[0] = []int{triangle[0][0]}

	for i := 1; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
		dp[i][0] = triangle[i][0] + dp[i-1][0]
	}

	for i := 1; i < len(triangle); i++ {
		for j := 1; j < len(triangle[i]); j++ {
			if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = MinInt(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	total := dp[len(triangle)-1][0]
	for i := 1; i < len(dp[len(triangle)-1]); i++ {
		total = MinInt(total, dp[len(triangle)-1][i])
	}

	return total
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// example 931
func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	sum := 0xffff
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = 0xffff
			if i-1 >= 0 {
				dp[i][j] = MinInt(dp[i-1][j]+matrix[i][j], dp[i][j])
				if j-1 >= 0 {
					dp[i][j] = MinInt(dp[i-1][j-1]+matrix[i][j], dp[i][j])
				}

				if j+1 < n {
					dp[i][j] = MinInt(dp[i-1][j+1]+matrix[i][j], dp[i][j])
				}
			} else {
				dp[i][j] = matrix[i][j]
			}

			if i == m-1 && sum > dp[i][j] {
				sum = dp[i][j]
			}
		}
	}
	return sum
}
