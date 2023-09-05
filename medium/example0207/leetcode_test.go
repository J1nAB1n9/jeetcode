package example0207

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(canFinish(2, [][]int{[]int{1, 0}}))
}

func TestExample2(t *testing.T) {
	fmt.Println(canFinish(2, [][]int{[]int{1, 0}, []int{0, 1}}))
}

func TestExample3(t *testing.T) {
	fmt.Println(canFinish(4, [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 0}, []int{3, 2}}))
}

func TestExample4(t *testing.T) {
	fmt.Println(canFinish(1, [][]int{}))
}

//  [[1,4],[2,4],[3,1],[3,2]]
func TestExample5(t *testing.T) {
	fmt.Println(canFinish(5, [][]int{[]int{1, 4}, []int{2, 4}, []int{3, 1}, []int{3, 2}}))
}
