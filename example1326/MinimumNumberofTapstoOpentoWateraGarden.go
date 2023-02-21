package example1326

import "sort"

type line struct {
	left   int
	right  int
	length int
}

type lines []line

func (s lines) Len() int      { return len(s) }
func (s lines) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s lines) Less(i, j int) bool {
	if s[i].left == s[j].left {
		return s[i].length > s[j].length
	}
	return s[i].left < s[j].left
}

func minTaps(n int, ranges []int) int {
	var arrLines lines
	for i, v := range ranges {
		l := line{}
		if v == 0 {
			continue
		}
		if v >= i {
			l.left = 0
		} else {
			l.left = i - v
		}

		if v+i >= n {
			l.right = n
		} else {
			l.right = v + i
		}

		l.length = l.right - l.left
		arrLines = append(arrLines, l)
	}
	if len(arrLines) == 0 {
		return -1
	}

	sort.Sort(arrLines)

	ans := 1
	lastArea := arrLines[0].right
	//for _, v := range arrLines {
	//	if v.left > lastArea {
	//		return -1
	//	}
	//	if v.right <= lastArea {
	//		continue
	//	}
	//
	//	lastArea = v.right
	//	ans++
	//}
	i := 0
	for i < len(arrLines)-1 {
		if arrLines[i].left > lastArea {
			return -1
		}

		j := i + 1
		k := j
		for ; j < len(arrLines); j++ {
			if arrLines[j].left > lastArea {
				break
			} else {
				if arrLines[j].right >= arrLines[k].right {
					k = j
				}
			}
		}
		i = k
		if lastArea >= arrLines[k].right {
			continue
		} else if lastArea < arrLines[k].left {
			return -1
		}
		ans++
		lastArea = arrLines[i].right
		if lastArea == n {
			return ans
		}
	}

	if lastArea != n {
		return -1
	}
	return ans
}
