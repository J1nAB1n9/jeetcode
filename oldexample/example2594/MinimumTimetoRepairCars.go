package example2594

import (
	"fmt"
)

func repairCars(ranks []int, cars int) int64 {
	size := len(ranks)
	maxx := int64(0)
	nr := make([]int, size)
	for i := 1; i <= cars; i++ {
		tm := 0
		index := 0
		for j := 0; j < size; j++ {
			tmj := (nr[j] + 1) * (nr[j] + 1) * ranks[j]
			if tm == 0 || tm >= tmj {
				tm = tmj
				index = j
			}
		}

		nr[index] += 1
		if int64(tm) > maxx {
			maxx = int64(tm)
		}
	}
	// PrintfArr(nr)
	return maxx
}

////////////////////////////////////////////////////////////////////
func PrintfArr(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
