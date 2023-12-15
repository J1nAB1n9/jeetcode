package example2415

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[5 : 7+1])
}

func TestExample1(t *testing.T) {
	reverseOddLevels(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   21,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   34,
				Left:  nil,
				Right: nil,
			},
		},
	})
}
