package example0007

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

func TestExample2(t *testing.T) {
	fmt.Println(reverse(120))
	fmt.Println(reverse(1000))
	fmt.Println(reverse(0))
}

func TestExample3(t *testing.T) {
	fmt.Println(reverse(2147483646))
}
