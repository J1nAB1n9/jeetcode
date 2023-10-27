package example1465

import "sort"

const mod = 1000000000 + 7

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	maxHighCutDis, maxWightCutDis := 0, 0
	for i, v := range horizontalCuts {
		pre := 0
		if i > 0 {
			pre = horizontalCuts[i-1]
		}

		if (v - pre) > maxHighCutDis {
			maxHighCutDis = v - pre
		}
	}
	if h-horizontalCuts[len(horizontalCuts)-1] > maxHighCutDis {
		maxHighCutDis = h - horizontalCuts[len(horizontalCuts)-1]
	}
	for i, v := range verticalCuts {
		pre := 0
		if i > 0 {
			pre = verticalCuts[i-1]
		}

		if (v - pre) > maxWightCutDis {
			maxWightCutDis = v - pre
		}
	}
	if w-verticalCuts[len(verticalCuts)-1] > maxWightCutDis {
		maxWightCutDis = w - verticalCuts[len(verticalCuts)-1]
	}

	return ((maxHighCutDis % mod) * (maxWightCutDis % mod)) % mod
}
