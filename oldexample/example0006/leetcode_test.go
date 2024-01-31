package example0006

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}

/*
PINALSIGYAHRPI
PINALSIGYAHRPI

P  I  N
AA SR G
Y LH I
P  I

P  I  N
AL SI G
Y AH R
P  I
*/
