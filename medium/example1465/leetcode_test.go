package example1465

import (
	"fmt"
	"testing"
)

//输入：h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
//输出：4
func TestExample1(t *testing.T) {
	fmt.Println(maxArea(5, 4, []int{1, 2, 4}, []int{1, 3}))
}

//输入：h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
//输出：6
func TestExample2(t *testing.T) {
	fmt.Println(maxArea(5, 4, []int{3, 1}, []int{1}))
}

//输入：h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
//输出：9
func TestExample3(t *testing.T) {
	fmt.Println(maxArea(5, 4, []int{3}, []int{3}))
}
