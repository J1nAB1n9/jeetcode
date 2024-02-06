package LCP0030

import (
	"container/heap"
)

type priorityQueue []int

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *priorityQueue) Pop() interface{} {
	val := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return val
}

func magicTower(nums []int) int {
	sum := 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	//fmt.Println(sum)
	if sum <= 0 {
		return -1
	}

	ans := 0
	// lastSum := 0 // 其实移到队尾的操作就是加到这里
	var negatives priorityQueue
	heap.Init(&negatives)
	negativeSum := 0
	sum = 1
	for _, v := range nums {
		if v >= 0 {
			sum += v
		} else {
			negativeSum += v
			heap.Push(&negatives, v)

			for sum+negativeSum <= 0 {
				ans += 1

				n := heap.Pop(&negatives).(int)
				negativeSum -= n
			}
		}
	}

	return ans
}
