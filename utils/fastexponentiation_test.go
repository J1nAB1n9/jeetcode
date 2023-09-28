package utils

import (
	"fmt"
	"math"
	"testing"
)

func TestFastExponentiation(t *testing.T) {
	fmt.Println(uint64(math.Pow(float64(2), float64(10))))
	fmt.Println(FastExponentiation(2, 10))
	fmt.Println("=========================")
	fmt.Println(math.Pow(float64(22), float64(11)))
	fmt.Println(FastExponentiation(22, 11))
	fmt.Println("=========================")
}
