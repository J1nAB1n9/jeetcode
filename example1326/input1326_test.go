package example1326

import (
	"fmt"
	"testing"
)

//5
//[3,4,1,1,0,0]

func TestDemo1(t *testing.T) {
	ranges := []int{3, 4, 1, 1, 0, 0}
	fmt.Println(minTaps(5, ranges))
}

//9
//[0,5,0,3,3,3,1,4,0,4]

func TestDemo2(t *testing.T) {
	ranges := []int{0, 5, 0, 3, 3, 3, 1, 4, 0, 4}
	fmt.Println(minTaps(9, ranges))
}

func TestDemo3(t *testing.T) {
	ranges := []int{0, 3, 3, 2, 2, 4, 2, 1, 5, 1, 0, 1, 2, 3, 0, 3, 1, 1}
	fmt.Println(minTaps(17, ranges))
}

func TestDemo4(t *testing.T) {
	ranges := []int{1, 0, 4, 0, 4, 1, 4, 3, 1, 1, 1, 2, 1, 4, 0, 3, 0, 3, 0, 3, 0, 5, 3, 0, 0, 1, 2, 1, 2, 4, 3, 0, 1, 0, 5, 2}
	fmt.Println(minTaps(35, ranges))
}

func TestDemo5(t *testing.T) {
	ranges := []int{1, 0, 0, 0, 1}
	fmt.Println(minTaps(4, ranges))
}
