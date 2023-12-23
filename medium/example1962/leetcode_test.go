package example1962

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Printf("ans:%d\n", minStoneSum([]int{5, 4, 9}, 2))
}
func TestExample2(t *testing.T) {
	fmt.Printf("ans:%d\n", minStoneSum([]int{4, 3, 6, 7}, 3))
}
func TestExample3(t *testing.T) {
	fmt.Printf("ans:%d\n", minStoneSum([]int{1}, 3))
}
