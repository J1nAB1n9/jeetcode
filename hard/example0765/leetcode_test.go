package example0765

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(minSwapsCouples([]int{0, 2, 1, 3}))
}

func TestExample2(t *testing.T) {
	fmt.Println(minSwapsCouples([]int{3, 2, 0, 1}))
}

func TestExample3(t *testing.T) {
	fmt.Println(minSwapsCouples([]int{1, 4, 0, 5, 8, 7, 6, 3, 2, 9}))
}
func TestExample4(t *testing.T) {
	fmt.Println(minSwapsCouples([]int{2, 0, 5, 4, 3, 1}))
}

func TestExample5(t *testing.T) {
	fmt.Println(minSwapsCouples([]int{6, 2, 1, 7, 4, 5, 3, 8, 0, 9}))
}
