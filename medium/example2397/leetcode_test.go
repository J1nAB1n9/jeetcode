package example2397

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2))
}
func TestExample2(t *testing.T) {
	fmt.Println(maximumRows([][]int{{1}, {0}}, 1))
}
func TestExample3(t *testing.T) {
	fmt.Println(maximumRows([][]int{{0}}, 1))
}
