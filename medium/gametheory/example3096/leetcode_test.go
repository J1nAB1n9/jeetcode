package example3096

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(minimumLevels([]int{1, 0, 1, 0}))
}

func TestExample2(t *testing.T) {
	fmt.Println(minimumLevels([]int{1, 1, 1, 1, 1}))
}

func TestExample3(t *testing.T) {
	fmt.Println(minimumLevels([]int{0, 0}))
}

func TestExample4(t *testing.T) {
	fmt.Println(minimumLevels([]int{1, 1}))
}
