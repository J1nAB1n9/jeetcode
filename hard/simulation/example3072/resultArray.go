package example3072

import (
	"math/rand"
	"time"
)

type node struct {
	key   int
	count int // 当前节点及其左子树的元素个数
	prior int // 优先级
	left  *node
	right *node
}

type treap struct {
	root *node
}

func newNode(key int) *node {
	return &node{
		key:   key,
		count: 1,
		prior: rand.Int(),
	}
}

func (t *treap) insert(key int) {
	t.root = t.insertRecursive(t.root, key)
}

func (t *treap) insertRecursive(node *node, key int) *node {
	if node == nil {
		return newNode(key)
	}

	if key < node.key {
		node.left = t.insertRecursive(node.left, key)
		if node.left.prior < node.prior {
			node = t.rotateRight(node)
		}
	} else {
		node.right = t.insertRecursive(node.right, key)
		if node.right.prior < node.prior {
			node = t.rotateLeft(node)
		}
	}

	node.count = 1 + countNodesInSubtree(node.left) + countNodesInSubtree(node.right)
	return node
}

func (t *treap) countGreaterThan(x int) int {
	return t.countGreaterThanRecursive(t.root, x)
}

func (t *treap) countGreaterThanRecursive(node *node, x int) int {
	if node == nil {
		return 0
	}

	if node.key > x {
		return 1 + countNodesInSubtree(node.right) + t.countGreaterThanRecursive(node.left, x)
	} else {
		return t.countGreaterThanRecursive(node.right, x)
	}
}

func (t *treap) rotateLeft(node *node) *node {
	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node
	newRoot.count = node.count
	node.count = 1 + countNodesInSubtree(node.left) + countNodesInSubtree(node.right)
	return newRoot
}

func (t *treap) rotateRight(node *node) *node {
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node
	newRoot.count = node.count
	node.count = 1 + countNodesInSubtree(node.left) + countNodesInSubtree(node.right)
	return newRoot
}

func countNodesInSubtree(node *node) int {
	if node == nil {
		return 0
	}
	return node.count
}

func resultArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	var arr1 = []int{nums[0]}
	var arr2 = []int{nums[1]}

	tree1 := &treap{}
	tree2 := &treap{}

	tree1.insert(nums[0])
	tree2.insert(nums[1])

	for i := 2; i < len(nums); i++ {
		val := nums[i]
		tree1Val := tree1.countGreaterThan(val)
		tree2Val := tree2.countGreaterThan(val)

		if tree1Val > tree2Val {
			arr1 = append(arr1, val)
			tree1.insert(val)
		} else if tree1Val < tree2Val {
			arr2 = append(arr2, val)
			tree2.insert(val)
		} else {
			if countNodesInSubtree(tree1.root) <= countNodesInSubtree(tree2.root) {
				arr1 = append(arr1, val)
				tree1.insert(val)
			} else {
				arr2 = append(arr2, val)
				tree2.insert(val)
			}
		}
	}

	arr1 = append(arr1, arr2...)
	return arr1
}
