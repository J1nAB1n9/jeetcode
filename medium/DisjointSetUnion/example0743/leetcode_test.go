package example0743

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
}

func TestExample2(t *testing.T) {
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1))
}
func TestExample3(t *testing.T) {
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 3, 7}, {1, 3, 4}, {2, 1, 2}}, 3, 2))
}
