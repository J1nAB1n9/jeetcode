package example1410

import (
	"fmt"
	"testing"
)

func TestArrCup(t *testing.T) {
	str := []rune("0123456789")
	fmt.Println(string(str[2:5]))
}

func TestExample1(t *testing.T) {
	fmt.Printf(entityParser("&amp; is an HTML entity but &ambassador; is not."))
}

func TestExample2(t *testing.T) {
	fmt.Printf(entityParser("and I quote: &quot;...&quot;"))
}

func TestExample3(t *testing.T) {
	fmt.Printf(entityParser("leetcode.com&frasl;problemset&frasl;all"))
}
