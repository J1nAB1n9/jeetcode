package example1962

import "container/heap"

type priorityQueue []int

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *priorityQueue) Pop() interface{} {
	val := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return val
}

func minStoneSum(piles []int, k int) int {
	sum := 0
	for _, v := range piles {
		sum += v
	}

	var queue priorityQueue
	queue = piles
	heap.Init(&queue)
	for i := 0; i < k && 0 < queue.Len(); i++ {
		val := heap.Pop(&queue).(int)
		sum -= val / 2
		val = (val + 1) / 2
		heap.Push(&queue, val)
	}

	return sum
}
