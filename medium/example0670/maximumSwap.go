package example0670

func maximumSwap(num int) int {
	//if num <= 10 {
	//	return num
	//}

	var numArr []int

	swapBefore := num
	for num != 0 {
		val := num % 10
		numArr = append(numArr, val)
		num = num / 10
	}

	ans := swapBefore
	in := 1
	for i := 0; i < len(numArr); i++ {
		jn := in * 10
		for j := i + 1; j < len(numArr); j++ {
			v, w := numArr[i], numArr[j]

			later := swapBefore - v*in + v*jn - w*jn + w*in
			if later > ans {
				ans = later
			}
			jn = jn * 10
		}
		in = in * 10
	}

	return ans
}

//
//func maximumSwap(num int) int {
//	num = oneSwap(num)
//	return oneSwap(num)
//}
