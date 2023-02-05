package example0011

func getMin(a, b int) int {
	if a >= b {
		return b
	}

	return a
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	area := 0
	n := len(height)
	i := 0
	j := n - 1
	for i < j {
		area = getMax(getMin(height[i], height[j])*(j-i), area)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return area
}
