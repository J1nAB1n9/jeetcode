package example2477

type node struct {
	son []*node
	val int
}

var ans int64

var mp map[int]struct{}

func dfs(n *node, seats int) (cars, num int64) {
	if n == nil {
		return
	}

	num = 1
	for _, v := range n.son {
		if _, ok := mp[v.val]; ok {
			continue
		}

		mp[v.val] = struct{}{}
		c, m := dfs(v, seats)
		cars += c
		num += m

		if m != 0 {
			ans += c + 1
		} else {
			ans += c
		}
	}

	if num >= int64(seats) {
		cars += num / int64(seats)
		num = num % int64(seats)
	}

	return
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	tree := make(map[int]*node)
	for _, v := range roads {
		l, r := v[0], v[1]
		if tree[l] == nil {
			tree[l] = &node{
				son: nil,
				val: l,
			}
		}

		if tree[r] == nil {
			tree[r] = &node{
				son: nil,
				val: r,
			}
		}

		nl, nr := tree[l], tree[r]
		nl.son = append(nl.son, nr)
		nr.son = append(nr.son, nl)
	}

	mp = make(map[int]struct{})
	mp[0] = struct{}{}

	ans = 0
	_, _ = dfs(tree[0], seats)

	//if m != 0 {
	//	ans += c + 1
	//} else {
	//	ans += c
	//}
	return ans
}
