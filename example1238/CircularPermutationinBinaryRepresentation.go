package example1238

var arr []int
var last int

func getNextSort(n int) {
	for j := 0; j < n; j++ {
		last = last ^ (1 << j)
		arr = append(arr, last)
		getNextSort(j)
	}
}

func circularPermutation(n int, start int) []int {
	arr = []int{}
	// 2*n - 1
	last = start
	arr = append(arr, last)
	for i := 0; i < n; i++ {
		last = last ^ (1 << i)
		arr = append(arr, last)
		getNextSort(i)
	}

	return arr
}

//func circularPermutation2(n int, start int) []int {
//	arr = []int{}
//	// 2*n - 1
//	arr = append(arr, start)
//	for i := 0; i < n; i++ {
//		start = start ^ (1 << i)
//		arr = append(arr, start)
//		for j := 0; j < i; j++ {
//			start = start ^ (1 << j)
//			arr = append(arr, start)
//		}
//	}
//
//	return arr
//}
