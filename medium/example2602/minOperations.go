package example2602

import (
	"fmt"
	"sort"
)

func abs(x int64) int64 {
	if x < 0 {
		x = -1 * x
	}
	return x
}
func minOperations(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	// test
	fmt.Println(nums)

	sumArr := make([]int64, len(nums))
	sumArr[0] = int64(nums[0])

	for i := 1; i < len(nums); i++ {
		sumArr[i] = sumArr[i-1] + int64(nums[i])
	}

	search := func(s int) int {
		l := 0
		r := len(nums) - 1
		for l < r {
			mid := (l + r) / 2
			if nums[mid] == s {
				return mid
			}
			if nums[mid] < s {
				l = mid + 1
				continue
			} else {
				r = mid
				continue
			}
		}

		return l
	}
	var ans []int64
	for _, q := range queries {
		qi := search(q)
		if nums[qi] > q {

		} else {
			v1 := abs(sumArr[qi] - int64((qi+1))*int64(q))
			v2 := abs(sumArr[qi]) - int64((qi+1))*int64(q))
		}

	}

	return ans
}
