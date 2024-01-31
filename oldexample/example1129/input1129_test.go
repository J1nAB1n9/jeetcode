package example1129

import (
	"fmt"
	"testing"
)

//Example 1:
//
//Input: n = 3, redEdges = [[0,1],[1,2]], blueEdges = []
//Output: [0,1,-1]
//
//Example 2:
//
//Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]
//Output: [0,1,-1]

func TestDemo1(t *testing.T) {
	ans := shortestAlternatingPaths(3, [][]int{[]int{0, 1}, []int{1, 2}}, [][]int{})
	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
	fmt.Println(" ")
}

func TestDemo2(t *testing.T) {
	ans := shortestAlternatingPaths(3, [][]int{[]int{0, 1}}, [][]int{[]int{2, 1}})
	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
	fmt.Println(" ")
}
