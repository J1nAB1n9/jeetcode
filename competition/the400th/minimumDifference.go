package the400th

import "math"

func minimumDifference(nums []int, k int) int {

	minDiff := math.MaxInt32

	return minDiff
}

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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
