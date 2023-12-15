package example2415

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[5 : 7+1])
}
