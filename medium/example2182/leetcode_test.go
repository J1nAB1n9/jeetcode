package example2182

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(repeatLimitedString("cczazcc", 3))
}

func TestExample2(t *testing.T) {
	fmt.Println(repeatLimitedString("aababab", 2))
}

func TestExample3(t *testing.T) {
	fmt.Println(repeatLimitedString("cczazcc", 2))
}
