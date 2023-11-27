package datastructure

import "fmt"

type SegmentTree struct {
	baseArr   []interface{}
	treeNodes []interface{}
	lazyData  []interface{}
	compare   func(interface{}, interface{}) interface{}
	update    func(interface{}, interface{}, int, int) interface{}
	arrayLen  int
}

// return a SegmentTree
func NewSegmentTree(base []interface{}, fn func(interface{}, interface{}) interface{}, fn2 func(interface{}, interface{}, int, int) interface{}) *SegmentTree {
	capacity := len(base)
	tree := &SegmentTree{
		baseArr:   base,
		treeNodes: make([]interface{}, 2*capacity+1),
		lazyData:  make([]interface{}, 2*capacity+1),
		compare:   fn,
		update:    fn2,
		arrayLen:  capacity,
	}

	// build tree
	tree.buildTree(0, 0, capacity-1)

	/////////////////////////////////////////
	//If you need to minimize memory
	//consumption as much as possible,
	//you can perform memory cleanup here.
	tree.baseArr = nil
	//////////////////////////////////////////
	return tree
}

func (tree *SegmentTree) buildTree(index int, left, right int) {
	if left == right {
		tree.treeNodes[index] = tree.baseArr[left]
		return
	}

	mid := (left + right) / 2
	tree.buildTree(index*2+1, left, mid)
	tree.buildTree((index+1)*2, mid+1, right)

	tree.treeNodes[index] = tree.compare(tree.treeNodes[index*2+1], tree.treeNodes[(index+1)*2])
}

func (tree *SegmentTree) SelectNode(_left, _right int /*Need Find*/) interface{} {
	return tree.selectNode(0, 0, tree.arrayLen-1, _left, _right)
}

func (tree *SegmentTree) selectNode(index int /*Index*/, left, right int /*Find Range*/, _left, _right int /*Need Find*/) interface{} {
	if left > _right || right < _left {
		return nil
	}

	if left >= _left && right <= _right {
		if tree.lazyData[index] == nil {
			return tree.treeNodes[index]
		}

		return tree.update(tree.treeNodes[index], tree.lazyData[index], left, right)
	}

	if tree.lazyData[index] != nil {
		tree.updateLazyData(index, left, right)
	}

	mid := (left + right) / 2
	return tree.compare(tree.selectNode(index*2+1, left, mid, _left, _right), tree.selectNode((index+1)*2, mid+1, right, _left, _right))
}

// Lazy Update 只适用于线性更新，例如：如果是每个节点平方就没辙了
func (tree *SegmentTree) UpdateNode(_left, _right int /*Need Find*/, updateVal interface{}) {
	tree.updateNode(0, 0, tree.arrayLen-1, _left, _right, updateVal)
}

// 向上更新
func (tree *SegmentTree) updateNode(index int /*Index*/, left, right int /*Find Range*/, _left, _right int /*Need Find*/, updateVal interface{}) interface{} {
	if left > _right || right < _left {
		if tree.lazyData[index] != nil {
			return tree.update(tree.treeNodes[index], tree.lazyData[index], left, right)
		}
		return tree.treeNodes[index]
	}

	if tree.lazyData[index] != nil {
		tree.updateLazyData(index, left, right)
	}

	if left >= _left && right <= _right {
		if tree.lazyData[index] == nil {
			tree.lazyData[index] = updateVal
		} else {
			tree.lazyData[index] = tree.update(tree.lazyData[index], updateVal, left, right)
		}
		return tree.update(tree.treeNodes[index], tree.lazyData[index], left, right)
	}

	mid := (left + right) / 2
	tree.treeNodes[index] = tree.compare(tree.updateNode(index*2+1, left, mid, _left, _right, updateVal), tree.updateNode((index+1)*2, mid+1, right, _left, _right, updateVal))
	return tree.treeNodes[index]
}

func (tree *SegmentTree) updateLazyData(index int, left, right int) {
	tree.treeNodes[index] = tree.update(tree.treeNodes[index], tree.lazyData[index], left, right)

	if index*2+1 < len(tree.treeNodes) {
		tree.lazyData[index*2+1] = tree.update(tree.lazyData[index*2+1], tree.lazyData[index], left, right)
	}

	if (index+1)*2 < len(tree.treeNodes) {
		tree.lazyData[(index+1)*2] = tree.update(tree.lazyData[(index+1)*2], tree.lazyData[index], left, right)
	}

	tree.lazyData[index] = nil
}

func (tree *SegmentTree) NaiveUpdate(_left, _right int /*Need Find*/, updateVal interface{}) {
	tree.naiveUpdate(0, 0, tree.arrayLen-1, _left, _right, updateVal)
}

// 无视所有的lazy
func (tree *SegmentTree) naiveUpdate(index int /*Index*/, left, right int /*Find Range*/, _left, _right int /*Need Find*/, updateVal interface{}) interface{} {
	if left > _right || right < _left {
		return tree.treeNodes[index]
	}

	if left == right {
		tree.treeNodes[index] = tree.update(tree.treeNodes[index], updateVal, left, right)
		return tree.treeNodes[index]
	}

	mid := (left + right) / 2
	tree.treeNodes[index] = tree.compare(tree.naiveUpdate(index*2+1, left, mid, _left, _right, updateVal), tree.naiveUpdate((index+1)*2, mid+1, right, _left, _right, updateVal))
	return tree.treeNodes[index]
}

func ShowTree(tree *SegmentTree, index int) {
	queue := make([]int, 0)
	queue = append(queue, index)
	queuePtl := make([]int, 0)
	for len(queue) != 0 {
		index = queue[0]
		queue = queue[1:]
		fmt.Printf(" %v:%v", tree.treeNodes[index], tree.lazyData[index])

		if index*2+1 < len(tree.treeNodes) {
			queuePtl = append(queuePtl, index*2+1)
		}

		if (index+1)*2 < len(tree.treeNodes) {
			queuePtl = append(queuePtl, (index+1)*2)
		}

		if len(queue) == 0 {
			fmt.Println()
			queue = queuePtl
			queuePtl = make([]int, 0)
		}
	}

}
