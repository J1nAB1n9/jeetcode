package example2982

import (
	"fmt"
	"testing"
)

func TestSubstrings(t *testing.T) {
	maximumLength("qqqqbbbbcc")
}

//2
func TestMaximumLength1(t *testing.T) {
	fmt.Println(maximumLength("aaaa"))
}

//-1
func TestMaximumLength2(t *testing.T) {
	fmt.Println(maximumLength("abcdef"))
}

//1
func TestMaximumLength3(t *testing.T) {
	fmt.Println(maximumLength("abcaba"))
}
