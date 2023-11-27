package datastructure

import (
	"fmt"
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	base := []interface{}{1, 2, 3, 4, 5}
	fn := func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	}
	fn2 := func(a, b interface{}, left, right int) interface{} {
		return a.(int) + b.(int)*(right-left+1)
	}
	tree := NewSegmentTree(base, fn, fn2)

	// 验证 SelectNode 方法
	result := tree.SelectNode(0, len(base)-1).(int)
	expectedResult := 15 // 1 + 2 + 3 + 4 + 5
	fmt.Printf("Expected: %d, Got: %d\n", expectedResult, result)

	// 验证 UpdateNode 方法
	tree.UpdateNode(2, 4, 10)
	ShowTree(tree, 0)
	result = tree.SelectNode(0, len(base)-1).(int)
	expectedResult = 45 // 1 + 2 + 3 + 4 + 5 + 10 + 10 + 10
	fmt.Printf("Expected: %d, Got: %d\n", expectedResult, result)
}

func TestSegmentTree_NaiveUpdate(t *testing.T) {
	base := []interface{}{1, 2, 3, 4, 5}
	fn := func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	}
	fn2 := func(a, b interface{}, left, right int) interface{} {
		return a.(int) * a.(int)
	}
	tree := NewSegmentTree(base, fn, fn2)

	// 验证 UpdateNode 方法
	tree.NaiveUpdate(2, 4, 0)
	ShowTree(tree, 0)
	result := tree.SelectNode(0, len(base)-1).(int)
	expectedResult := 53 // 1 + 2 + 3*3 + 4*4 + 5*5
	fmt.Printf("Expected: %d, Got: %d\n", expectedResult, result)
}
