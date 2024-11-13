package example3261

type node struct {
	l, r int
}

func (n node) right(x, y int) bool {
	if x <= n.l && y >= n.r {
		return false
	}
	return true
}

func (n node) area(x, y int) int64 {
	if n.l == x && n.r == y {
		return int64(1)
	} else if n.l == x || n.r == y {
		return int64(getAbs(n.l-x) + getAbs(y-n.r) + 1)
	}

	return int64(getAbs(n.l-x) + getAbs(y-n.r) + 2)

}

func getAbs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func getArrSum(size int) int64 {
	n := int64(size)
	return n * (n + 1) / 2
}

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	size := len(s)

	var nodes []node
	limit := -1

	for r := 0; r < size; r++ {
		num0 := 0
		num1 := 0
		if s[r] == '0' {
			num0++
		} else {
			num1++
		}

		for l := r - 1; l > limit; l-- {
			if s[l] == '0' {
				num0++
			} else {
				num1++
			}

			if num0 > k && num1 > k {
				limit = l
				nodes = append(nodes, node{
					l: l,
					r: r,
				})
				break
			}
		}
	}
	// fmt.Println(nodes)

	ansList := make([]int64, len(queries), len(queries))
	for i, q := range queries {
		x := q[0]
		y := q[1]

		last := -1
		ansList[i] = getArrSum(y - x + 1)
		for _, n := range nodes {
			if !n.right(x, y) {
				if last == -1 {
					ansList[i] -= n.area(x, y)
				} else {
					ansList[i] -= n.area(last+1, y)
				}
				last = n.l
			}
		}
	}

	return ansList
}
