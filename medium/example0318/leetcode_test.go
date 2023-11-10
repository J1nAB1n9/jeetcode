package example0318

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestExample(t *testing.T) {
	var i int
	size := unsafe.Sizeof(i)
	fmt.Printf("int size on 64-bit Windows: %d bytes\n", size)
}

func TestExample1(t *testing.T) {
	fmt.Printf("ans:%d\n", maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

func TestExample2(t *testing.T) {
	fmt.Printf("ans:%d\n", maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
}

func TestExample3(t *testing.T) {
	fmt.Printf("ans:%d\n", maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
}
