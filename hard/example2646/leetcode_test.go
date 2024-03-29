package example2646

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(minimumTotalPrice(9, [][]int{{2, 5}, {3, 4}, {4, 1}, {1, 7}, {6, 7}, {7, 0}, {0, 5}, {5, 8}}, []int{4, 4, 6, 4, 2, 4, 2, 14, 8}, [][]int{{1, 5}, {2, 7}, {4, 3}, {1, 8}, {2, 8}, {4, 3}, {1, 5}, {1, 4}, {2, 1}, {6, 0}, {0, 7}, {8, 6}, {4, 0}, {7, 5}, {7, 5}, {6, 0}, {5, 1}, {1, 1}, {7, 5}, {1, 7}, {8, 7}, {2, 3}, {4, 1}, {3, 5}, {2, 5}, {3, 7}, {0, 1}, {5, 8}, {5, 3}, {5, 2}}))
}

func TestExample2(t *testing.T) {
	fmt.Println(minimumTotalPrice(9, [][]int{{2, 5}, {3, 4}, {4, 1}, {1, 7}, {6, 7}, {7, 0}, {0, 5}, {5, 8}}, []int{4, 4, 6, 4, 2, 4, 2, 14, 8}, [][]int{{2, 7}}))
}
