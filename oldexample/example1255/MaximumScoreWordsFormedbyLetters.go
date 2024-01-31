package example1255

//var mapLatter map[int]int
//var mapScore map[int]int
//
//func GetScoreWord(word string) int {
//	score := 0
//	useNem := make(map[int]int)
//	for _, c := range word {
//		useNem[int(c)] = useNem[int(c)] + 1
//		if useNem[int(c)] > mapLatter[int(c)] {
//			return 0
//		}
//
//		if useNem[int(c)] == 1 {
//			score += mapScore[int(c)]
//		}
//
//	}
//
//	return score
//}

func maxScoreWords(words []string, letters []byte, score []int) int {
	ans := 0
	letterNum := make([]int, 26)
	for _, l := range letters {
		letterNum[l-'a'] += 1
	}
	len := len(words)
	for i := 1; i <= (1 << len); i++ {
		nowScore := 0
		worldNum := make([]int, 26)
		f := true
		for j, w := range words {
			if (1<<j)&i == 0 {
				continue
			}

			for _, c := range w {
				cv := c - 'a'
				worldNum[cv]++
				nowScore += score[cv]
				if worldNum[cv] > letterNum[cv] {
					f = false
					break
				}
			}
		}

		if !f {
			continue
		}
		if ans < nowScore {
			ans = nowScore
		}
	}
	return ans
}
