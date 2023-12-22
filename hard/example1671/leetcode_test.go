package example1671

import (
	"fmt"
	"testing"
)

//func TestRangeMap(t *testing.T) {
//	mp := make(map[int]int)
//	index := 7
//	for i := 1; i < index; i++ {
//		mp[i] = i
//	}
//
//	for k, v := range mp {
//		mp[index] = index
//		index = index + 1
//		fmt.Printf("%d %d\n", k, v)
//		delete(mp, k+1)
//	}
//
//	fmt.Printf("len:%d index:%d\n", len(mp), index)
//}

func TestExample1(t *testing.T) {
	fmt.Println(minimumMountainRemovals([]int{1, 3, 1}))
}

func TestExample2(t *testing.T) {
	fmt.Println(minimumMountainRemovals([]int{2, 1, 1, 5, 6, 2, 3, 1}))
}
