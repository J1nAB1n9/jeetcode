package example1944

func canSeePersonsCount(heights []int) (ans []int) {
	ans = make([]int, len(heights))
	size := len(heights)
	stack := []int{size - 1}
	for i := size - 2; i >= 0; i-- {
		minStackNum := stack[len(stack)-1]
		if heights[i] < heights[minStackNum] {
			stack = append(stack, i)
			ans[i] = 1
		} else {
			popNum := 0
			popIndex := stack[len(stack)-1]
			for len(stack) != 0 && heights[i] > heights[popIndex] {
				stack = stack[:len(stack)-1]
				popNum++
				if len(stack) == 0 {
					break
				}
				popIndex = stack[len(stack)-1]
			}

			if len(stack) == 0 {
				ans[i] = popNum
			} else {
				ans[i] = popNum + 1
			}
			stack = append(stack, i)
		}
	}
	return
}
