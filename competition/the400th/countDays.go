package the400th

import "sort"

type meets [][]int

func (s meets) Len() int      { return len(s) }
func (s meets) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s meets) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}
	return s[i][0] < s[j][0]
}
func countDays(days int, meetings [][]int) int {
	mt := meets(meetings)
	sort.Sort(mt)

	ans := 0
	begin := 1
	end := days
	for i := 0; i < len(mt); i++ {
		if begin < mt[i][0] {
			if i == 0 {
				ans += mt[i][0] - begin
			} else {
				ans += mt[i][0] - begin - 1
			}

			begin = mt[i][1]
		} else if begin < mt[i][1] {
			begin = mt[i][1]
		}
	}
	if begin < end {
		ans += end - begin
	}

	return ans
}
