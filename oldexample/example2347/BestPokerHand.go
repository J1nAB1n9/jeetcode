package example2347

func bestHand(ranks []int, suits []byte) string {
	a := suits[0]
	find := true
	for _, v := range suits {
		if v != a {
			find = false
			break
		}
	}
	if find {
		return "Flush"
	}
	maxx := 0
	torank := make(map[int]int)
	for _, v := range ranks {
		torank[v] += 1
		if torank[v] >= maxx {
			maxx = torank[v]
		}
	}

	if maxx >= 3 {
		return "Three of a Kind"
	} else if maxx >= 2 {
		return "Pair"
	}

	return "High Card"
}
