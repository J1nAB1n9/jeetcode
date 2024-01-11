package example2645

func addMinimum(word string) int {
	// 'a','b','c'
	ans := 0
	runeWord := []rune(word)
	b := runeWord[0]
	if b == 'b' {
		ans += 1
	} else if b == 'c' {
		ans += 2
	}

	for i := 1; i < len(runeWord); i++ {
		r := runeWord[i]
		if b == 'a' {
			if r == 'a' {
				ans += 2
			} else if r == 'b' {
				//continue
			} else if r == 'c' {
				ans += 1
			}
		} else if b == 'b' {
			if r == 'a' {
				ans += 1
			} else if r == 'b' {
				ans += 2
			} else if r == 'c' {
				//continue
			}
		} else if b == 'c' {
			if r == 'a' {
				//continue
			} else if r == 'b' {
				ans += 1
			} else if r == 'c' {
				ans += 2
			}
		}
		b = r

	}

	if b == 'a' {
		ans += 2
	} else if b == 'b' {
		ans += 1
	}
	return ans
}

func AddMinimum(word string) int {
	return addMinimum(word)
}
