package example0670

import (
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println(maximumSwap(123456))
}

func TestExample2(t *testing.T) {
	//7236
	fmt.Println(maximumSwap(2736))
}

func TestExample3(t *testing.T) {
	fmt.Println(maximumSwap(9973))
}

func TestExample4(t *testing.T) {
	fmt.Println(maximumSwap(1))
}

func TestExample5(t *testing.T) {
	fmt.Println(maximumSwap(18))
	fmt.Println(maximumSwap(10))
}

func TestExample6(t *testing.T) {
	fmt.Println(maximumSwap(100))
}

func TestExample7(t *testing.T) {
	fmt.Println(maximumSwap(20))
}
