package the400th

func minimumChairs(s string) int {
	stack := 0
	ans := 0
	for _, v := range s {
		if v == 'E' {
			stack += 1
		} else {
			stack -= 1
		}
		if stack >= ans {
			ans = stack
		}
	}
	return ans
}
