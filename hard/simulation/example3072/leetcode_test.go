package example3072

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(resultArray([]int{2, 1, 3, 3}))
}

func TestExample2(t *testing.T) {
	fmt.Println(resultArray([]int{5, 14, 3, 1, 2}))
}

func TestNewTree(t *testing.T) {
	tree := &treap{}
	tree.insert(5)
	tree.insert(3)
	tree.insert(1)

	fmt.Print(tree.countGreaterThan(2))
}
