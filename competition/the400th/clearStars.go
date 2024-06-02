package the400th

import (
	"container/heap"
	"fmt"
	"strings"
)

const nd = '*'

// 定义一个Item类型，它是优先队列中的元素类型
type Item struct {
	value    int  // 元素值
	priority rune // 优先级
	index    int  // 元素在堆中的索引
}

// 定义一个ItemHeap类型，它代表一个最小堆
type ItemHeap []*Item

// 实现heap.Interface中的Len方法
func (h ItemHeap) Len() int { return len(h) }

// 实现heap.Interface中的Less方法
func (h ItemHeap) Less(i, j int) bool {
	if h[i].priority == h[j].priority {
		return h[i].value > h[j].value
	}

	return h[i].priority < h[j].priority
}

// 实现heap.Interface中的Swap方法
func (h ItemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

// 实现heap.Interface中的Push方法
func (h *ItemHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.index = n
	*h = append(*h, item)
}

// 实现heap.Interface中的Pop方法
func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // 避免内存泄漏
	item.index = -1 // 表示这个item已经被移除了
	*h = old[0 : n-1]
	return item
}
func clearStars(s string) string {
	need := make(map[int]struct{})
	pq := make(ItemHeap, 0)
	heap.Init(&pq)
	ar := []rune(s)
	for i, v := range ar {
		if v == nd {
			item := heap.Pop(&pq).(*Item)
			need[item.value] = struct{}{}
			need[i] = struct{}{}
			continue
		}

		item := &Item{
			value:    i,
			priority: v,
		}
		heap.Push(&pq, item)
	}

	// fmt.Println(need)
	for k, v := range need {
		fmt.Printf("key:%d,value:%d\n", k, v)
	}

	buff := strings.Builder{}
	for i, v := range ar {
		if _, ok := need[i]; ok {
			continue
		}

		buff.WriteRune(v)
	}

	return buff.String()
}
