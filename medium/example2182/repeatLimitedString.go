package example2182

import (
	"bytes"
	"strings"
)

type table struct {
	c rune
	n int
}

type stack struct {
	data []table
}

func (s stack) Len() int {
	return len(s.data)
}

func (s stack) Empty() bool {
	return s.Len() == 0
}

func (s *stack) Push(t table) {
	s.data = append(s.data, t)
}

func (s *stack) Pop() table {
	t := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return t
}

func (s *stack) Top() rune {
	if s.Len() == 0 {
		return 0
	}

	r := s.data[s.Len()-1].c
	s.data[s.Len()-1].n -= 1
	if s.data[s.Len()-1].n == 0 {
		s.Pop()
	}
	return r
}

func repeatLimitedString(s string, repeatLimit int) string {
	var ts stack
	mp := make(map[int]int)
	for _, c := range s {
		index := int(c - 'a')
		mp[index] += 1
	}

	for i := 0; i < 26; i++ {
		if v, has := mp[i]; has {
			t := table{
				c: rune('a' + i),
				n: v,
			}
			ts.Push(t)
		}
	}

	ans := strings.Builder{}
	for !ts.Empty() {
		t := ts.Pop()
		standard := bytes.Buffer{}
		if t.n >= repeatLimit {
			for i := 0; i < repeatLimit; i++ {
				standard.WriteRune(t.c)
			}
		}

		for t.n > 0 {
			if t.n >= repeatLimit {
				t.n -= repeatLimit
				ans.Write(standard.Bytes())

				if t.n != 0 {
					next := ts.Top()
					if next == 0 {
						break
					}
					ans.WriteRune(next)
				}
			} else {
				for ; t.n > 0; t.n-- {
					ans.WriteRune(t.c)
				}
			}
		}
	}

	return ans.String()
}
