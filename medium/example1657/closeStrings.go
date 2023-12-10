package example1657

import "sort"

func closeStrings(word1 string, word2 string) bool {
	size := len(word1)
	if size != len(word2) {
		return false
	}

	var ra rune
	ra = 'a'

	arr1 := make([]int, 26)
	arr2 := make([]int, 26)

	for _, v := range word1 {
		arr1[v-ra] += 1
	}
	for _, v := range word2 {
		arr2[v-ra] += 1
	}

	for i := 0; i < 26; i++ {
		if arr1[i]*arr2[i] == 0 && arr1[i]+arr2[i] != 0 {
			return false
		}
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
