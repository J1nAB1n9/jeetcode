package example0064

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	if len(grid) < 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = MinInt(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
		//fmt.Println(dp[i])
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。

//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
