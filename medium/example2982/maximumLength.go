package example2982

import (
	"fmt"
	"sort"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type specialSubstring struct {
	r rune
	n int
}

type substrings []specialSubstring

func (s substrings) Len() int      { return len(s) }
func (s substrings) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s substrings) Less(i, j int) bool {
	if s[i].n == s[j].n {
		return s[i].r < s[j].r
	}
	return s[i].n > s[j].n
}

func (s substrings) ToString() {
	for _, v := range s {
		fmt.Printf("char:%s num:%d\n", string(v.r), v.n)
	}
}
func maximumLength(s string) int {
	//ss := substrings{}
	_ss := make([]substrings, 26)
	ptl := specialSubstring{}
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if ptl.r != r {
			if ptl.n != 0 {
				// ss = append(ss, ptl)
				index := ptl.r - 'a'
				_ss[index] = append(_ss[index], ptl)
			}
			ptl.r = r
			ptl.n = 1
		} else {
			ptl.n += 1
		}
	}
	if ptl.n != 0 {
		index := ptl.r - 'a'
		_ss[index] = append(_ss[index], ptl)
	}

	for _, v := range _ss {
		if len(v) == 0 {
			continue
		}

		sort.Sort(v)
	}

	// sort.Sort(ss)

	ans := -1
	for _, ss := range _ss {
		for i, s := range ss {
			if ans >= s.n {
				// 这种情况肯定是已经满足
				return ans
			}

			if s.n >= 3 {
				ans = getMax(ans, s.n-2)
			} else if s.n == 2 && i+1 < len(ss) {
				if ss[i+1].r == s.r {
					ans = getMax(ans, 1)
				}
			}

			if i+2 < len(ss) {
				if s.r == ss[i+1].r && ss[i+1].r == ss[i+2].r {
					if s.n == ss[i+1].n && ss[i+1].n == ss[i+2].n {
						ans = getMax(ans, s.n)
					} else {
						mini := getMin(getMin(s.n, ss[i+1].n), ss[i+2].n)
						ans = getMax(ans, mini)
					}
				}
			}
		}
	}

	// ss.ToString()
	return ans
}
