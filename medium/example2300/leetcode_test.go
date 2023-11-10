package example2300

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(successfulPairs([]int{1, 2, 3}, []int{1, 2, 6, 7, 9}, 5))
}

func TestExample2(t *testing.T) {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}

func TestExample3(t *testing.T) {
	fmt.Println(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
}
