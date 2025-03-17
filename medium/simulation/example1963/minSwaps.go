package example1963

func minSwaps(s string) int {
	stack := make([]rune, 0, len(s))
	for _, r := range s {
		if r == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}

	if len(stack) == 0 {
		return 0
	}

	swapCount := len(stack) / 4
	if len(stack)%4 != 0 {
		swapCount += 1
	}
	return swapCount
}
