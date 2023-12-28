package example1954

const table_size = 62996

var table []int64

func init() {
	table = make([]int64, table_size+1)
	for n := int64(1); n <= int64(table_size); n++ {
		a1 := (int64(1) + n) * n / int64(2)
		an := a1 + n*n
		table[n] = (a1 + an) * (n + int64(1)) * int64(2)
	}
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
