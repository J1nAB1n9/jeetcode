package example2379
func Min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumRecolors(blocks string, k int) int {
	strSize := len(blocks)
	pd := make([]int,strSize)
	// 计算前缀和
	for i,v := range blocks {
		if i != 0 {
			pd[i] = pd[i-1]
		}

		if v == int32('B') {
			pd[i] += 1
		}
	}

	// 计算结果
	ans := k
	for i := 0;i + k -1 < strSize;i++ {
		j := i + k - 1
		lineNum := pd[j] - pd[i]
		if blocks[i] == 'B' {
			lineNum += 1
		}

		if lineNum >= k {
			return 0
		}

		ans = Min(ans,k - lineNum)
	}
	return ans
}