package example2258

import (
	"fmt"
	"testing"
)

// [[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]
func TestExample1(t *testing.T) {
	fmt.Println(maximumMinutes([][]int{{0, 2, 0, 0, 0, 0, 0}, {0, 0, 0, 2, 2, 1, 0}, {0, 2, 0, 0, 1, 2, 0}, {0, 0, 2, 2, 2, 0, 2}, {0, 0, 0, 0, 0, 0, 0}}))
}
func TestExample2(t *testing.T) {
	fmt.Println(maximumMinutes([][]int{{0, 0, 0, 0}, {0, 1, 2, 0}, {0, 2, 0, 0}}))
}
func TestExample3(t *testing.T) {
	fmt.Println(maximumMinutes([][]int{{0, 0, 0}, {2, 2, 0}, {1, 2, 0}}))
}

//[[10,48],[43,50],[35,35],[38,39]]
//[49,1,32,15,35,16,37,44,10,49]
