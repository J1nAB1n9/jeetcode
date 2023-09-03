package example1921

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(eliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1}))
}

func TestExample2(t *testing.T) {
	fmt.Println(eliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1}))
}

//输入：dist = [3,2,4], speed = [5,3,2]
//输出：1

func TestExample3(t *testing.T) {
	fmt.Println(eliminateMaximum([]int{3, 2, 4}, []int{5, 3, 2}))
}
