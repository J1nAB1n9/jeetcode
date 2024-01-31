package example0011

import (
	"fmt"
	"testing"
)

func TestDemo1(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(input))
}

func TestDemo2(t *testing.T) {
	input := []int{1, 1}
	fmt.Println(maxArea(input))
}

func TestDemo3(t *testing.T) {
	input := []int{1, 2, 1}
	fmt.Println(maxArea(input))
}
