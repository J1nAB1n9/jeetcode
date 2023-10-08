package datastructure

import "fmt"

func FastExponentiation(base uint64, exp int) (uint64, error) {
	if exp > 64 {
		return 0, fmt.Errorf("pow > 64")
	}
	result := uint64(1)
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}

		base *= base
		exp >>= 1
	}
	return result, nil
}
