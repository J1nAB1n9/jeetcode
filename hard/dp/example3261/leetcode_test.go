package example3261

import (
	"fmt"
	"testing"
)

// 15 9 3
// 26
func TestCountKConstraintSubstrings(t *testing.T) {
	fmt.Println(countKConstraintSubstrings("010101", 1, [][]int{{0, 5}, {1, 4}, {2, 3}}))
	fmt.Println(countKConstraintSubstrings("0001111", 2, [][]int{{0, 6}}))
}
