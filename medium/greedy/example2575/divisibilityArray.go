package example2575

func divisibilityArray(word string, m int) []int {
	n := len(word)
	arr := make([]int, n)

	w := 0
	for i, v := range word {
		w = (w*10 + int(v-'0')) % m
		if w == 0 {
			arr[i] = 1
		}
	}

	return arr
}
