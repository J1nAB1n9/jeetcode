package example2300

import "sort"

//func successfulPairs(spells []int, potions []int, success int64) []int {
//	sort.Ints(potions)
//	ans := make([]int, len(spells))
//	for i := 0; i < len(spells); i++ {
//		minPotions := int(success / int64(spells[i]))
//		if int64(spells[i]) > success {
//			minPotions = 0 // 处理特殊情况
//		} else if int64(minPotions)*int64(spells[i]) != success {
//			minPotions += 1
//		}
//
//		l, r := 0, len(potions)-1
//
//		index := -1
//		// mid select
//		for l <= r {
//			mid := (l + r) / 2
//			if potions[mid] == minPotions {
//				index = mid
//				break
//			} else if potions[mid] > minPotions {
//				r = mid - 1
//			} else {
//				l = mid + 1
//			}
//		}
//
//		if index != -1 {
//			ans[i] = len(potions) - index
//		} else {
//			if l >= len(potions) {
//				ans[i] = 0
//			} else if r==-1||l==-1{
//				ans[i] = len(potions)
//			}else if potions[r] > minPotions {
//				ans[i] = len(potions) - r
//			} else if potions[l] > minPotions {
//				ans[i] = len(potions) - l
//			} else {
//				ans[i] = 0
//			}
//		}
//

//	}
//	return ans
//}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	suffix := make(map[int]int)
	for i, v := range potions {
		if _, has := suffix[v]; has == false {
			suffix[v] = i
		}
	}

	ans := make([]int, len(spells))
	for i := 0; i < len(spells); i++ {
		minPotions := int(success / int64(spells[i]))
		if int64(spells[i]) > success {
			minPotions = 0 // 处理特殊情况
		} else if int64(minPotions)*int64(spells[i]) != success {
			minPotions += 1
		}

		l, r := 0, len(potions)-1

		index := -1
		// mid select
		for l <= r {
			mid := (l + r) / 2
			if potions[mid] == minPotions {
				index = mid
				break
			} else if potions[mid] > minPotions {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if index != -1 {
			index := suffix[potions[index]]
			ans[i] = len(potions) - index
		} else {
			if l >= len(potions) {
				ans[i] = 0
			} else if r == -1 || l == -1 {
				ans[i] = len(potions)
			} else if potions[r] > minPotions {
				index := suffix[potions[r]]
				ans[i] = len(potions) - index
			} else if potions[l] > minPotions {
				index := suffix[potions[l]]
				ans[i] = len(potions) - index
			} else {
				ans[i] = 0
			}
		}

	}
	return ans
}
