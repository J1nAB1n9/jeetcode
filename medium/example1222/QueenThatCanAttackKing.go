package example1222

type node struct {
	x, y int
}

var ms map[node]struct{}
var ans [][]int

func while(x, y int, v, w int) {
	for true {
		x += v
		y += w

		if x >= 8 || y >= 8 || x < 0 || y < 0 {
			break
		}

		n := node{x: x, y: y}
		if _, ok := ms[n]; ok {
			ans = append(ans, []int{x, y})
			return
		}
	}

}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	ms = make(map[node]struct{})
	ans = make([][]int, 0, len(queens))
	for _, v := range queens {
		n := node{
			x: v[0],
			y: v[1],
		}

		ms[n] = struct{}{}
	}

	while(king[0], king[1], 0, 1)
	while(king[0], king[1], 1, 1)
	while(king[0], king[1], 1, 0)
	while(king[0], king[1], 0, -1)
	while(king[0], king[1], -1, -1)
	while(king[0], king[1], -1, 0)
	while(king[0], king[1], -1, 1)
	while(king[0], king[1], 1, -1)

	return ans
}
