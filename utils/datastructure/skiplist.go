package datastructure

import (
	"fmt"
	"math/rand"
	"time"
)

type SkipListIntegerNode struct {
	key, value int // key,val
	isHead     bool
	down       *SkipListIntegerNode // 下一层节点
	next       *SkipListIntegerNode // 下一个节点
}

type SkipListInteger struct {
	maxLevel int // 跳表最高层级
	curLevel int // 当前层级

	head *SkipListIntegerNode   // 头节点
	list []*SkipListIntegerNode // 所有节点(无序)
}

func NewSkipListInteger(maxLevel int) SkipListInteger {
	return SkipListInteger{
		maxLevel: maxLevel,
		curLevel: 0,
		head: &SkipListIntegerNode{
			key:    -1,
			value:  -1,
			isHead: true,
			down:   nil,
			next:   nil,
		},
		list: []*SkipListIntegerNode{},
	}
}

func (sl *SkipListInteger) randomLvl() int {
	rand.Seed(time.Now().UnixNano())
	level := 1
	for rand.Intn(10) < 5 && level < sl.maxLevel {
		level += 1
	}
	return level
}

func (sl *SkipListInteger) Insert(key, val int) {
	level := sl.randomLvl()

	// 假如层高较低，拉高层高
	for sl.curLevel < level {
		nHead := &SkipListIntegerNode{
			key:    -1,
			value:  -1,
			isHead: true,
			down:   sl.head,
			next:   nil,
		}

		sl.head = nHead
		sl.curLevel += 1
	}

	// 只需要在同一层级搜索就行
	node := sl.head
	for l := sl.curLevel; l > level; l-- {
		node = node.down
	}

	var preNode *SkipListIntegerNode
	for node != nil {
		if node.next == nil || node.next.key > key {
			// find = true
			newNode := &SkipListIntegerNode{
				key:   key,
				value: val,
				down:  nil,
				next:  node.next,
			}

			if preNode != nil {
				preNode.down = newNode
			}
			preNode = newNode

			node.next = newNode
			node = node.down
			continue
		}
		node = node.next
	}
}

func (sl SkipListInteger) SelectValueByKey(key int) (int, bool) {
	node := sl.head

	for node != nil {
		if node.key == key && !node.isHead {
			return node.value, true
		}

		if node.next == nil || node.next.key > key {
			node = node.down
			continue
		}

		node = node.next
	}

	return -1, false
}

func (sl *SkipListInteger) PrintfInfo() {
	head := sl.head
	for i := 0; i < sl.curLevel && head != nil; i++ {
		fmt.Printf("head%d:	", i+1)
		p := head
		for p.next != nil {
			fmt.Printf("	kv[%d,%d]", p.next.key, p.next.value)
			p = p.next
		}
		head = head.down
		fmt.Println()
	}

}
