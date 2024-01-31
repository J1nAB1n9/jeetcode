package example2602

import (
	"fmt"
	"testing"
)

//输入：nums = [3,1,6,8], queries = [1,5]
//输出：[14,10]
//
//输入：nums = [2,9,6,3], queries = [10]
//输出：[20]

func TestExample1(t *testing.T) {
	fmt.Println(minOperations([]int{3, 1, 6, 8}, []int{1, 5}))
}

func TestExample2(t *testing.T) {
	fmt.Println(minOperations([]int{2, 9, 6, 3}, []int{10}))
}
