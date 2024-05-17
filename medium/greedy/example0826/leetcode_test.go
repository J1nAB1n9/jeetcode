package example0826

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
}

func TestExample2(t *testing.T) {
	fmt.Println(maxProfitAssignment([]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}))
}
