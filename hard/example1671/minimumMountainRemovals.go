package example1671

import "fmt"

func getMax(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func minimumMountainRemovals(nums []int) int {
	dp1 := make([]int, len(nums))
	dp2 := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp1[i] = getMax(dp1[i], dp1[j]+1)
			}
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j >= i; j-- {
			if nums[i] > nums[j] {
				dp2[i] = getMax(dp2[i], dp2[j]+1)
			}
		}
	}

	fmt.Printf("dp1:%v\n", dp1)
	fmt.Printf("dp2:%v\n", dp2)

	ans := 0
	for i := 1; i < len(nums)-1; i++ {
		ans = getMax(ans, dp1[i]+dp2[i]+1)
	}

	return len(nums) - ans
}
