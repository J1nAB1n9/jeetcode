package example0740

import (
	"fmt"
	"testing"
)

func TestDeleteAndEarn(t *testing.T) {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))          // 6
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4})) // 9
}
