package example1123

import (
	"fmt"
	"testing"
)

func Ptff(n *TreeNode) {
	fmt.Println(n.Val)
	fmt.Printf("%v %v\n", n.Left, n.Right)
}

func TestExample(t *testing.T) {
	node := lcaDeepestLeaves(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	})

	Ptff(node)
}

func TestExample2(t *testing.T) {
	node := lcaDeepestLeaves(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	})
	Ptff(node)
}
