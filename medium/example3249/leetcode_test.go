package example3249

import (
	"fmt"
	"testing"
)

//{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}}

func Test_co(t *testing.T) {
	fmt.Println(countGoodNodes([][]int{{2, 0}, {4, 2}, {1, 2}, {3, 1}, {5, 1}}))
}

func Test_co2(t *testing.T) {
	fmt.Println(countGoodNodes([][]int{{1, 0}, {3, 0}, {2, 3}}))
}

func Test_co3(t *testing.T) {
	fmt.Println(countGoodNodes([][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {0, 5}, {5, 6}, {6, 7}, {7, 8}, {0, 9}, {9, 10}, {9, 12}, {10, 11}}))
}
