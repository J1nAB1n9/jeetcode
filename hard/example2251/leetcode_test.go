package example2251

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(fullBloomFlowers([][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}, []int{2, 3, 7, 11}))
}

func TestExample2(t *testing.T) {
	fmt.Println(fullBloomFlowers([][]int{{1, 10}, {3, 3}}, []int{3, 3, 2}))
}

//[[10,48],[43,50],[35,35],[38,39]]
//[49,1,32,15,35,16,37,44,10,49]
