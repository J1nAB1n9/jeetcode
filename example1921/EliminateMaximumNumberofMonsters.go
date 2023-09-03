package example1921

import "sort"

func eliminateMaximum(dist []int, speed []int) int {
	monstersNum := len(dist)
	landingTime := make([]float64, monstersNum)

	for i := 0; i < monstersNum; i++ {
		landingTime[i] = float64(dist[i]) / float64(speed[i])
	}

	sort.Float64s(landingTime)
	ans := float64(0.0)
	i := 0
	for i < monstersNum && ans < landingTime[i] {
		i++
		ans += float64(1.0)
	}
	return int(ans)
}
