package example2594

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := make([]int, 10)
	for _, v := range slice {
		fmt.Println(v)
	}
}

func TestExample(t *testing.T) {
	fmt.Println(repairCars([]int{4, 2, 3, 1}, 10))
}

func TestExample2(t *testing.T) {
	fmt.Println(repairCars([]int{5, 1, 8}, 6))
}
