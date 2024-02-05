package example1696

import (
	"fmt"
	"math"
)

// 10^5
type line struct {
	l int
	r int
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxResult(nums []int, k int) int {
	ans := nums[0]
	if len(nums) == 1 {
		return ans
	}

	var ls []line
	// 正数必拿
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			l := line{
				l: i,
			}

			size := 1
			for i = i + 1; i < len(nums); i++ {
				if nums[i] >= 0 {
					break
				}
				size++
			}

			if size >= k {
				l.r = i - 1
				ls = append(ls, l)
			}
		}

		if i < len(nums) {
			ans += nums[i]
		}
	}

	fmt.Println(ls)
	isend := false
	for _, l := range ls {
		size := l.r - l.l + 1
		dp := make([]int, size)
		// k步以内的肯定走一次就够了
		var i int
		for i = 0; i < k && i < size; i++ {
			dp[i] = nums[i+l.l]
		}

		// 超过k步的只能dp了
		for ; i < size; i++ {
			val := nums[i+l.l]
			minStep := math.MinInt64
			for j := i - 1; j >= i-k; j-- {
				if val+dp[j] > minStep {
					minStep = val + dp[j]
				}
			}
			dp[i] = minStep
		}

		maxx := math.MinInt64
		if l.r == len(nums)-1 {
			maxx = dp[size-1]
			isend = true
		} else {
			for z := 0; z < k; z++ {
				if dp[size-1-z] > maxx {
					maxx = dp[size-1-z]
				}
			}
		}
		ans += maxx
	}

	if isend == false && nums[len(nums)-1] < 0 {
		ans += nums[len(nums)-1]
	}
	return ans
}

func maxResult2(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	queue := make([]int, n) // 模拟双端队列
	qi, qj := 0, 1
	for i := 1; i < n; i++ {
		for qi < qj && queue[qi] < i-k {
			qi++
		}
		dp[i] = dp[queue[qi]] + nums[i]
		for qi < qj && dp[queue[qj-1]] <= dp[i] {
			qj--
		}
		queue[qj] = i
		qj++
	}
	return dp[n-1]

}
