package example0907

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
}
func TestExample2(t *testing.T) {
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}
