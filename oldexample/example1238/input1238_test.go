package example1238

import (
	"fmt"
	"testing"
)

//5
//[3,4,1,1,0,0]
//func getNextSort(n int, start int) {
//	for j := 0; j < n; j++ {
//		start = start ^ (1 << j)
//		arr = append(arr, start)
//		getNextSort(j, start)
//	}
//}[2 3 1 0 5 4 6 7]

func TestDemo1(t *testing.T) {
	fmt.Println(circularPermutation(2, 3))
}

// 0010 0011 0001 0000 0100 0101 0111 0110
// [2 3 1 0 4 5 7]
func TestDemo2(t *testing.T) {
	fmt.Println(circularPermutation(3, 2))
	//fmt.Println(circularPermutation2(3, 2))
}
