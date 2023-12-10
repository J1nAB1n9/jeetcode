package example2661

type node struct {
	x, y int
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	mp := make(map[int]node)
	ls, rs := 0, 0
	ls = len(mat)
	for i, k := range mat {
		rs = len(k)
		for j, v := range k {
			n := node{
				x: i,
				y: j,
			}
			mp[v] = n
		}
	}

	ml, mr := make(map[int]int), make(map[int]int)
	for i, v := range arr {
		ml[mp[v].x] += 1
		mr[mp[v].y] += 1

		if ml[mp[v].x] == ls || mr[mp[v].y] == rs {
			return i
		}
	}
	return 0
}
