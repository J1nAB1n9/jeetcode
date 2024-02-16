package example0107

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(levelOrderBottom(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}

func TestExample2(t *testing.T) {
	fmt.Println(levelOrderBottom(&TreeNode{
		Val: 1,
	}))
}
func TestExample3(t *testing.T) {
	var p *TreeNode
	fmt.Println(levelOrderBottom(p))
}
