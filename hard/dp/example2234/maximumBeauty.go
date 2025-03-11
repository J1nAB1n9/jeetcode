package example2234

import "sort"

//type Garden struct {
//	flowerCount int64
//	count       int64
//}
//type Gardens []Garden
//
//func (g Gardens) Len() int           { return len(g) }
//func (g Gardens) Less(i, j int) bool { return g[i].flowerCount > g[j].flowerCount }
//func (g Gardens) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }

func getMin(a, b int64) int64 {
	if a >= b {
		return b
	}
	return a
}

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {

	sort.Slice(flowers, func(i, j int) bool {
		return flowers[i] > flowers[j]
	})

	length := len(flowers)

	if flowers[length-1] >= target {
		return int64(length) * int64(full)
	}

	index := length - 1
	suffixSums := make([]int64, length)
	for i := length - 2; i >= 0; i-- {
		if index != i+1 || flowers[i] >= target {
			suffixSums[i] = 0xffffffff
			continue
		}
		suffixSums[i] = suffixSums[i+1] + int64(flowers[i]-flowers[i+1])*int64(length-i-1)
		if suffixSums[i] <= newFlowers {
			index = i
		}
	}

	cnt := int64(0)
	partialCnt := int64(flowers[index])
	if (newFlowers - suffixSums[index]) >= int64(length-index) {
		partialCnt = getMin((newFlowers-suffixSums[index])/int64(length-index)+int64(flowers[index]), int64(target-1))
	}
	beautyScore := partialCnt * int64(partial)
	for i := 0; i < length; i++ {
		if flowers[i] >= target {
			cnt++
			beautyScore += int64(full)
			continue
		}

		if int64(target-flowers[i]) > newFlowers {
			break
		}
		cnt++
		newFlowers -= int64(target - flowers[i])

		if index <= i {
			index = i + 1
		}

		for ; index < length-1 && suffixSums[index] > newFlowers; index++ {

		}

		partialCnt = int64(0)
		if index < length {
			if (newFlowers - suffixSums[index]) >= int64(length-index) {
				partialCnt = getMin((newFlowers-suffixSums[index])/int64(length-index)+int64(flowers[index]), int64(target-1))
			} else {
				partialCnt = int64(flowers[index])
			}
		}

		if cnt*int64(full)+partialCnt*int64(partial) > beautyScore {
			beautyScore = cnt*int64(full) + partialCnt*int64(partial)
		}
	}

	return beautyScore
}
