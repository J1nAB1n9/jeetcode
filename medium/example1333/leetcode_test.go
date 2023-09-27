package example1333

import (
	"fmt"
	"testing"
)

//todo: go test -v
//=== RUN   TestExample1
//[3 1 5]
//--- PASS: TestExample1 (0.00s)
//=== RUN   TestExample2
//[4 3 2 1 5]
//--- PASS: TestExample2 (0.00s)
//=== RUN   TestExample3
//[4 5]
//--- PASS: TestExample3 (0.00s)
//PASS
//ok      jeetcode/medium/example1333     1.251s

func TestExample1(t *testing.T) {
	fmt.Println(filterRestaurants([][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 1, 50, 10))
}

func TestExample2(t *testing.T) {
	fmt.Println(filterRestaurants([][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 0, 50, 10))
}

func TestExample3(t *testing.T) {
	fmt.Println(filterRestaurants([][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}, 0, 20, 3))
}
