package example0174

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(calculateMinimumHP([][]int{[]int{-2, -3, 3}, []int{-5, -10, 1}, []int{10, 30, -5}}))
}

func TestExample2(t *testing.T) {
	fmt.Println(calculateMinimumHP([][]int{[]int{0}}))
}
