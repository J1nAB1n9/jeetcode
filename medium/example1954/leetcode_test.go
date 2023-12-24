package example1954

import (
	"fmt"
	"testing"
)

func getABS(x int64) int64 {
	if x < 0 {
		return -1 * x
	}
	return x
}
func TestPrint(t *testing.T) {
	var arr []int64
	for i := 1; i <= 100; i++ {
		left := int64(-i)
		right := int64(i)
		ans := int64(0)
		for j := left; j <= right; j++ {
			for k := left; k <= right; k++ {
				ans += getABS(j) + getABS(k)
			}
		}
		arr = append(arr, ans)
	}
	//	z := (1 + i) * i / 2
	//	z2 :=
	//	arr2 = append(arr2, )
	//}
	//
	fmt.Println(arr)
}

func TestGetTable(t *testing.T) {
	var maxv int64
	maxv = 1000000000000000
	var arr []int64
	for n := int64(1); n <= int64(100000); n++ {
		a1 := (1 + n) * n / 2
		an := a1 + n*n
		ans := (a1 + an) * (n + 1) * 2
		if ans > maxv {
			arr = append(arr, (a1+an)*(n+1)*2)
			break
		}
		arr = append(arr, (a1+an)*(n+1)*2)
	}
	// fmt.Println(arr)
	fmt.Println(len(arr), arr[len(arr)-1], arr[len(arr)-2])
}

//时间	0ms 	击败 100.00%使用 Go 的用户
//内存	2.35MB	击败 8.00%使用 Go 的用户
func TestExample1(t *testing.T) {
	fmt.Printf("%d\n", minimumPerimeter(7642190260848))
}
