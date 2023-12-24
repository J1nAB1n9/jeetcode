package example1954

import "fmt"

const table_size = 62996

var table []int64

func init() {
	// todo: del Code
	f := false
	table = make([]int64, table_size+1)

	for n := int64(1); n <= int64(table_size); n++ {
		a1 := (int64(1) + n) * n / int64(2)
		an := a1 + n*n
		table[n] = (a1 + an) * (n + int64(1)) * int64(2)
		// todo: del Code
		if f == false && table[n] > 7642190260848 {
			f = true
			fmt.Printf("index:%v value:%v\n", n, table[n])
		}
	}

	// todo: del Code
	fmt.Printf("index:%v value:%v\n", 1551, table[1551])
	fmt.Printf("index:%v value:%v\n", 12410, table[12410])
	fmt.Printf("index:%v value:%v\n", 12409, table[12409])

}

func minimumPerimeter(neededApples int64) int64 {
	if neededApples == 0 {
		return 0
	}

	var left, right, mid int64
	left, right = 1, table_size
	mid = (left + right) / 2
	for left < right {
		if table[mid] == neededApples {
			return mid
		}

		if table[mid] > neededApples {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	if table[mid] < neededApples {
		mid = mid + 1
	}
	return mid * 8
}
