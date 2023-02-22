package example1138

import (
	"fmt"
	"testing"
)

//示例 1：
//输入：target = "leet"
//输出："DDR!UURRR!!DDD!"
//
//示例 2：
//输入：target = "code"
//输出："RR!DDRR!UUL!R!"

func TestDemo1(t *testing.T) {
	fmt.Println(alphabetBoardPath("leet"))
	fmt.Println(alphabetBoardPath("code"))
}

func TestDemo2(t *testing.T) {
	fmt.Println(alphabetBoardPath("zz"))
}
