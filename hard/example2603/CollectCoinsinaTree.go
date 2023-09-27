package example2603

const k = 2

type node struct {
	degree int
	next   []int
}

var tree []node

func clearTree(queue []int, coins []int, ans int, ik bool) int {
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		tree[q].degree--
		ans--

		for _, v := range tree[q].next {
			tree[v].degree--
			if ik {
				if tree[v].degree == 1 && coins[v] == 0 {
					queue = append(queue, v)
				}
			}

		}
	}

	return ans
}

func noCoinsNode() {

}

func collectTheCoins(coins []int, edges [][]int) int {

	size := len(coins)
	tree = make([]node, size)

	for _, v := range edges {
		x := v[0]
		y := v[1]

		tree[x].next = append(tree[x].next, y)
		tree[x].degree++
		tree[y].next = append(tree[y].next, x)
		tree[y].degree++
	}

	ans := size
	queue := make([]int, 0)
	for i, v := range tree {
		if v.degree == 1 && coins[i] == 0 {
			queue = append(queue, i)
		}
	}

	ans = clearTree(queue, coins, ans, true)

	queue = []int{}
	for i := 0; i < k; i++ {
		queue = []int{}
		for i, v := range tree {
			if v.degree == 1 {
				queue = append(queue, i)
			}
		}

		ans = clearTree(queue, coins, ans, false)
	}

	if ans == 0 {
		return 0
	}
	return (ans - 1) * 2

}
