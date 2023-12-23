package example0041

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}

func TestExample2(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
func TestExample3(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
