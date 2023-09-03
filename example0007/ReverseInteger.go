package example0007

import "fmt"

var maxNum string

func init() {
	maxNum = fmt.Sprintf("%d", 0x7FFFFFFF)
}

func reverse(x int) int {
	var zs bool
	size := len(maxNum)
	if x < 0 {
		zs = false
		x = -1 * x
	} else {
		zs = true
	}

	var ans []rune
	i := 0
	for x > 0 {
		c := x % 10
		// 过滤前缀0
		if len(ans) == 0 && c == 0 {
			x = x / 10
			continue
		}
		ans = append(ans, rune(c+'0'))
		x = x / 10

		i++

		if i > size {
			return 0
		}
	}

	fmt.Println(string(ans))
	if i >= size && maxNum < string(ans) {
		return 0
	}
	x = 0
	for _, v := range ans {
		c := int(v - '0')
		x = x*10 + c
	}

	if zs != true {
		return -1 * x
	}
	return x
}
