package example2477

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(minimumFuelCost([][]int{{0, 1}, {0, 2}, {0, 3}}, 5))
}

func TestExample2(t *testing.T) {
	fmt.Println(minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2))
}

func TestExample3(t *testing.T) {
	fmt.Println(minimumFuelCost([][]int{}, 1))
}
