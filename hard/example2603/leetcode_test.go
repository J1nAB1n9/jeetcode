package example2603

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(collectTheCoins([]int{1, 0, 0, 0, 0, 1}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}
