package example3096

func minimumLevels(possible []int) int {
	length := len(possible)
	prefixSum := make([]int, length, length)
	prefixSum[0] = possible[0]
	if prefixSum[0] == 0 {
		prefixSum[0] = -1
	}

	for i := 1; i < length; i++ {
		if possible[i] == 0 {
			possible[i] = -1
		}
		prefixSum[i] = prefixSum[i-1] + possible[i]
	}

	sumNum := prefixSum[length-1]

	for i, v := range prefixSum {
		if i == length-1 {
			break
		}

		if v > sumNum-prefixSum[i] {
			return i + 1
		}
	}
	return -1
}
