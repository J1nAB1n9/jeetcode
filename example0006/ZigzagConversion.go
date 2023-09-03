package example0006

import "strings"

/*
1  7
2 68
35 8
4  10


*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	zu := numRows*2 - 2
	ss := make([][]byte, numRows)

	for i, c := range s {
		n := (i + 1) % zu
		if n == 0 {
			n = zu
		}

		if n > numRows {
			n = 2*numRows - n
		}

		ss[n-1] = append(ss[n-1], byte(c))
	}

	var sb strings.Builder
	for _, s := range ss {
		sb.Write(s)
	}
	return sb.String()
}
