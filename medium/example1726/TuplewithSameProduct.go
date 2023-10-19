package example1726

func tupleSameProduct(nums []int) int {
	size := len(nums)
	hash := make(map[int]int)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			hash[nums[i]*nums[j]]++
		}
	}

	ans := 0
	for _, v := range hash {
		ans += v * (v - 1) * 4
	}
	return ans
}
