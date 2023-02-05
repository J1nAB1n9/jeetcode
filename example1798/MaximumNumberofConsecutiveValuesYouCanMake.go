package example1798

import "sort"

type gSet map[int]int

//// 这个方法感觉太慢了,而且有bug
//func getMaximumConsecutive(coins []int) int {
//	maxNumConsecutive := 1
//	exitMap := make(gSet)
//	minLimit := coins[0]
//	maxLimit := 0
//	// 初始化map,使得map中的每个key都指向比他大的值
//	sort.Ints(coins)
//	last := coins[0]
//
//	for _, val := range coins {
//		if last < val {
//			exitMap[last] = val
//		}
//
//		if val < minLimit {
//			minLimit = val
//		}
//
//		maxLimit += val
//		exitMap[val] = val
//	}
//
//	// 循环min-max之间的数值
//	for n := minLimit; n <= maxLimit; n = exitMap[n] {
//		if n == exitMap[n] {
//			break
//		}
//
//		// 判断n+1是否存在
//		if index, ok := exitMap[n+1]; ok {
//			maxNumConsecutive++
//		} else if {
//			maxNumConsecutive = 1
//		}
//	}
//
//	return maxNumConsecutive
//}

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	maxMumConsecutive := 1
	for _, val := range coins {
		if val > maxMumConsecutive {
			break
		}
		maxMumConsecutive += val
	}

	return maxMumConsecutive
}
