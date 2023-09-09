package example0210

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(findOrder(2, [][]int{[]int{0, 1}, []int{1, 0}}))
}

func TestExample2(t *testing.T) {
	fmt.Println(findOrder(2, [][]int{[]int{1, 0}}))
}

func TestExample3(t *testing.T) {
	fmt.Println(findOrder(1, [][]int{}))
}

func TestExample4(t *testing.T) {
	fmt.Println(findOrder(3, [][]int{[]int{1, 0}}))
}
