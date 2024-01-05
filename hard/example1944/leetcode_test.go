package example1944

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
}

func TestExample2(t *testing.T) {
	fmt.Println(canSeePersonsCount([]int{5, 1, 2, 3, 10}))
}
