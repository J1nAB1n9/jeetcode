package example1138

import "strings"

type node struct {
	x int
	y int
}
type Board map[byte]node

func (n node) Equip(o node) bool {
	if n.x == o.x && n.y == o.y {
		return true
	}

	return false
}

func getBoard() *Board {
	board := make(Board)
	for i := 'a'; i <= 'z'; i++ {
		index := int(i) - int('a')
		n := node{
			x: index / 5,
			y: index % 5,
		}
		board[byte(i)] = n
	}

	return &board
}

func mabs(x int) int {
	if x < 0 {
		x = x * -1
	}
	return x
}

func alphabetBoardPath(target string) string {
	var strBuilder strings.Builder
	board := getBoard()
	lastNode := node{
		x: 0,
		y: 0,
	}
	for _, v := range target {
		vnode := (*board)[byte(v)]
		if vnode.Equip(lastNode) {
			strBuilder.WriteString("!")
			continue
		}

		if lastNode.x == 5 && lastNode.y == 0 {
			strBuilder.WriteString("U")
			lastNode.x = 4
		}

		z := false
		if vnode.x == 5 && vnode.y == 0 {
			z = true
			vnode.x = 4
		}

		xy := mabs(vnode.y - lastNode.y)
		for i := 0; i < xy; i++ {
			if vnode.y > lastNode.y {
				strBuilder.WriteString("R")
			} else {
				strBuilder.WriteString("L")
			}
		}

		xy = mabs(vnode.x - lastNode.x)
		for i := 0; i < xy; i++ {
			if vnode.x > lastNode.x {
				strBuilder.WriteString("D")
			} else {
				strBuilder.WriteString("U")
			}
		}

		lastNode = vnode
		if z {
			strBuilder.WriteString("D")
			lastNode.x = 5
		}

		strBuilder.WriteString("!")
	}

	return strBuilder.String()
}
