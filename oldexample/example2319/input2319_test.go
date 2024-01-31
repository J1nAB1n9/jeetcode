package example2319

import (
	"fmt"
	"testing"
)

//grid = [[2,0,0,1],[0,3,1,0],[0,5,2,0],[4,0,0,2]]
//grid = [[5,7,0],[0,3,1],[0,5,0]]
func TestDemo1(t *testing.T) {
	var input [][]int
	input = append(input, []int{2, 0, 0, 1})
	input = append(input, []int{0, 3, 1, 0})
	input = append(input, []int{0, 5, 2, 0})
	input = append(input, []int{4, 0, 0, 2})

	if checkXMatrix(input) {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
}
