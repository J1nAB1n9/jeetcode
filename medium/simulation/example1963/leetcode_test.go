package example1963

import (
	"fmt"
	"testing"
)

func TestMinSwaps(t *testing.T) {
	fmt.Println(minSwaps("][]["))
	fmt.Println(minSwaps("]]][[["))
	fmt.Println(minSwaps("[]"))
}
