package LCP0030

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(magicTower([]int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}))
}

func TestExample2(t *testing.T) {
	fmt.Println(magicTower([]int{-200, -300, 400, 0}))
}
