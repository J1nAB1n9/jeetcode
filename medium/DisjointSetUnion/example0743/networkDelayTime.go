package example0743

import "container/heap"

type treeNode struct {
	index int
	cost  int
}

type treeHeap []treeNode

func (h treeHeap) Len() int           { return len(h) }
func (h treeHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h treeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *treeHeap) Push(x interface{}) {
	*h = append(*h, x.(treeNode))
}
func (h *treeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1:n]
	return x
}

type node struct {
	next []int
	cost []int
}

func networkDelayTime(times [][]int, n int, k int) int {
	nodes := make([]node, n+1, n+1)
	for _, t := range times {
		nodes[t[0]].next = append(nodes[t[0]].next, t[1])
		nodes[t[0]].cost = append(nodes[t[0]].cost, t[2])
	}

	maxCost := 0
	mp := make(map[int]int)
	h := &treeHeap{treeNode{
		index: k,
		cost:  0,
	}}
	mp[k] = 0
	heap.Init(h)

	for h.Len() > 0 {
		tn := h.Pop().(treeNode)

		if tn.cost > mp[tn.index] {
			continue
		}

		for j, m := range nodes[tn.index].next {
			nowCost := tn.cost + nodes[tn.index].cost[j]
			if nc, ok := mp[m]; ok {
				if nc <= nowCost {
					continue
				}
			}

			mp[m] = nowCost
			h.Push(treeNode{
				index: m,
				cost:  nowCost,
			})
		}
	}

	if len(mp) != n {
		return -1
	}

	for _, cost := range mp {
		if maxCost < cost {
			maxCost = cost
		}
	}
	return maxCost
}
