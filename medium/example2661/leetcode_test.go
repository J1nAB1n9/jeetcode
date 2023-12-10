package example2661

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(firstCompleteIndex([]int{2, 7, 4, 5, 6, 9}, [][]int{{2, 5}, {4, 6}, {7, 9}}))
}
