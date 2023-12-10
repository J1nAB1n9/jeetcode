package example1094

func carPooling(trips [][]int, capacity int) bool {
	arr := make([]int, 1005)
	np := 0
	for _, v := range trips {
		nm, from, to := v[0], v[1], v[2]
		arr[from] += nm
		arr[to] -= nm
	}

	for i := 0; i <= 1000; i++ {
		if arr[i] == 0 {
			continue
		}

		np += arr[i]
		if np > capacity {
			return false
		}
	}

	return true
}
