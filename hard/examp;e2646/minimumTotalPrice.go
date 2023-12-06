package examp_e2646

type node struct {
	pre  *node
	sons []*node
	val  int
	high int
}

var mp map[int]struct{}

func (n *node) buildTree(pre *node, h int) {
	n.high = h
	n.pre = pre
	for _, v := range n.sons {
		if _, has := mp[v.val]; has {
			continue
		}

		mp[v.val] = struct{}{}
		v.buildTree(n, n.high+1)
	}
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	tree := make(map[int]*node)
	for _, v := range edges {
		n1, n2 := tree[v[0]], tree[v[1]]

		if n1 == nil {
			n1 = &node{
				val: v[0],
			}
			tree[v[0]] = n1
		}

		if n2 == nil {
			n2 = &node{
				val: v[1],
			}
			tree[v[1]] = n2
		}

		n1.sons = append(n1.sons, n2)
		n2.sons = append(n2.sons, n1)
	}

	root := tree[0]
	if root == nil {
		return 0
	}

	// 建树
	mp = make(map[int]struct{})
	mp[0] = struct{}{}
	root.buildTree(nil, 0)

	// anso, anst := make([]int, len(trips)), make([]int, len(trips))
	anso, anst := 0, 0
	for _, v := range trips {
		l, r := v[0], v[1]
		n1, n2 := tree[l], tree[r]
		for n1.val != n2.val {
			h := 0
			onPrice := 0
			if n1.high == n2.high {
				h = n1.high
				onPrice = price[n1.val] + price[n2.val]
				n1 = n1.pre
				n2 = n2.pre
			} else if n1.high < n2.high {
				h = n2.high
				onPrice = price[n2.val]
				n2 = n2.pre
			} else {
				h = n1.high
				onPrice = price[n1.val]
				n1 = n1.pre
			}

			if h%2 == 0 {
				anso += onPrice
			} else {
				anst += onPrice
			}
		}

	}

	if anso > anst {
		return anso/2 + anst
	}

	return anso + anst/2
}
