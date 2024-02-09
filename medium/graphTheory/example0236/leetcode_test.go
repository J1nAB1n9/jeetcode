package example0236

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(lowestCommonAncestor(&TreeNode{
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
	}, &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}))
}

func TestExample2(t *testing.T) {
	fmt.Println(lowestCommonAncestor(&TreeNode{
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
	}, &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}))
}
func TestExample3(t *testing.T) {
	fmt.Println(lowestCommonAncestor(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}, &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}))
}
