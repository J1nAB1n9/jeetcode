package the400th

import (
	"fmt"
	"testing"
)

func TestDemo1(t *testing.T) {
	fmt.Printf("a:%v *:%v\n", rune('a'), rune('*'))
}

func TestCountDays1(t *testing.T) {
	fmt.Println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}}))
}

func TestCountDays2(t *testing.T) {
	fmt.Println(countDays(8, [][]int{{3, 4}, {4, 8}, {2, 5}, {3, 8}}))
}
