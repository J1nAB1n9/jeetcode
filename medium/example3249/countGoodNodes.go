package example3249

import "fmt"

type node struct {
	index int
	score int
	next  []*node
}

var ans int

func (n *node) f(i int) int {
	if len(n.next) <= 1 && n.index != 0 {
		//n.score = 1
		ans += 1
	} else {
		equip := true
		last := -1
		for _, q := range n.next {
			if q.index == i {
				continue
			}
			n.score += q.f(n.index)
			if last == -1 {
				last = q.score
			} else if last != q.score {
				equip = false
			}
		}

		if equip {
			ans += 1
		}
	}

	return n.score
}
func countGoodNodes(edges [][]int) int {
	nodes := make([]*node, len(edges)+1, len(edges)+1)
	createNode := func(index int) *node {
		if nodes[index] == nil {
			nodes[index] = &node{
				index: index,
				score: 1,
				next:  nil,
			}
		}
		return nodes[index]
	}

	// create tree
	for _, edge := range edges {
		left := createNode(edge[0])
		right := createNode(edge[1])

		left.next = append(left.next, right)
		right.next = append(right.next, left)
	}
	ans = 0
	// 统计节点
	nodes[0].f(-1)
	for _, v := range nodes {
		fmt.Printf("%d %d\n", v.index, v.score)
	}
	return ans
}

// 直接用数组会快很多
func countGoodNodes2(edges [][]int) int {
	tree := make([][]int, len(edges)+1)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}
	ans := 0
	var dfs func(int, int) int
	dfs = func(node int, father int) int {
		if len(tree[node]) == 0 {
			ans++
			return 1
		}
		cnt := 1
		pre := -1
		isGood := true
		for _, child := range tree[node] {
			if child == father {
				continue
			}
			cntChild := dfs(child, node)
			if pre == -1 {
				pre = cntChild
			} else if cntChild != pre {
				isGood = false
			}
			cnt += cntChild
		}
		if isGood {
			ans++
		}
		return cnt
	}
	dfs(0, -1)
	return ans
}
