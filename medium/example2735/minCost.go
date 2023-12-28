package example2735

func minElement(nums []int) int {
	k := 0
	for i := range nums {
		if nums[k] > nums[i] {
			k = i
		}
	}
	return k
}

func minElementInt64(nums []int64) int {
	k := 0
	for i := range nums {
		if nums[k] > nums[i] {
			k = i
		}
	}
	return k
}

func partialSum(F []int64) {
	for i := 1; i < len(F); i++ {
		F[i] += F[i-1]
	}
}

func minCost(nums []int, x int) int64 {
	n := len(nums)
	// 找出 nums 中最小的元素，并用其为首元素构造一个新的数组
	minIdx := minElement(nums)
	var tmp []int
	for i := 0; i < n; i++ {
		tmp = append(tmp, nums[(minIdx+i)%n])
	}
	nums = tmp

	L, R := make([]int, n), make([]int, n)
	L[0] = n - 1
	// 循环来看，右侧 nums[0] 是更小的元素，但不一定是第一个更小的元素，需要用单调栈计算得到
	for i := 0; i < n; i++ {
		R[i] = n - i - 1
	}
	s := []int{0}
	for i := 1; i < n; i++ {
		for len(s) > 0 && nums[i] < nums[s[len(s)-1]] {
			R[s[len(s)-1]] = i - s[len(s)-1] - 1
			s = s[:len(s)-1]
		}
		L[i] = i - s[len(s)-1] - 1
		s = append(s, i)
	}

	F := make([]int64, n)
	// 辅助函数，一次差分，将 F[l..r] 都增加 d
	diffOnce := func(l, r int, d int64) {
		if l > r {
			return
		}
		if l < n {
			F[l] += d
		}
		if r+1 < n {
			F[r+1] -= d
		}
	}
	// 辅助函数，二次差分，将 F[l..r] 增加 ki + b，i 是下标
	diffTwice := func(l, r int, k, b int64) {
		if l > r {
			return
		}
		diffOnce(l, l, k*int64(l)+b)
		diffOnce(l+1, r, k)
		diffOnce(r+1, r+1, -(k*int64(r) + b))
	}

	// 进行操作需要的成本
	diffTwice(0, n-1, int64(x), 0)

	for i := 0; i < n; i++ {
		minV := min(L[i], R[i])
		maxV := max(L[i], R[i])
		// 第一种情况，窗口数量 k+1，总和 nums[i] * k + nums[i]
		diffTwice(0, minV, int64(nums[i]), int64(nums[i]))
		// 第二种情况，窗口数量 minV+1，总和 0 * k + nums[i] * (minV + 1)
		diffTwice(minV+1, maxV, 0, int64(nums[i])*int64(minV+1))
		// 第三种情况，窗口数量 L[i]+R[i]-k+1，总和 -nums[i] * k + nums[i] * (L[i] + R[i] + 1)
		diffTwice(maxV+1, L[i]+R[i], -int64(nums[i]), int64(nums[i])*int64(L[i]+R[i]+1))
	}

	// 计算两次前缀和
	for i := 0; i < 2; i++ {
		partialSum(F)
	}
	return F[minElementInt64(F)]
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}
func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
