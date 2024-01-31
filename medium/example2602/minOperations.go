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
	ans := make([]int64, len(queries))
	for i, q := range queries {
		if nums[0] > q || q > nums[len(nums)-1] {
			ans[i] = abs(sumArr[len(nums)-1] - int64(q*(len(nums))))
			continue
		}

		qi := search(q)
		if nums[qi] > q {
			v1 := abs(sumArr[qi-1] - int64(qi)*int64(q))
			v2 := abs((sumArr[len(nums)-1] - sumArr[qi-1]) - int64(len(nums)-qi)*int64(q))
			ans[i] = v1 + v2
		} else {
			v1 := abs(sumArr[qi] - int64((qi+1))*int64(q))
			v2 := abs((sumArr[len(nums)-1] - sumArr[qi]) - int64(len(nums)-(qi+1))*int64(q))
			ans[i] = v1 + v2
		}
	}
	return ans
}
