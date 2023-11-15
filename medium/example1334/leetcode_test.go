package example1334

import (
	"fmt"
	"testing"
)

//todo: go test -v
//输入：n = 4, edges = {{0,1,3},{1,2,1},{1,3,4},{2,3,1}}, distanceThreshold = 4
//输出：3
//
//输入：n = 5, edges = {{0,1,2},{0,4,8},{1,2,3},{1,4,2},{2,3,1},{3,4,1}}, distanceThreshold = 2
//输出：0

func TestExample1(t *testing.T) {
	fmt.Println(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4))
}

func TestExample2(t *testing.T) {
	fmt.Println(findTheCity(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2))
}
