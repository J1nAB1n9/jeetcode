package example0064

import (
	"fmt"
	"leetcode/medium/dp/example0062"
	"testing"
)

func TestMinInt(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func TestMinInt2(t *testing.T) {
	fmt.Println(example0062.U([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
