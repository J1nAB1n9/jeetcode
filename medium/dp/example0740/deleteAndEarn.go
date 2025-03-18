package example0740

import "sort"

type number struct {
	val   int
	score int
}

type numbers []number

func (n numbers) Len() int           { return len(n) }
func (n numbers) Less(i, j int) bool { return n[i].val < n[j].val }
func (n numbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deleteAndEarn(nums []int) int {
	mp := make(map[int]int)
	for _, n := range nums {
		if _, ok := mp[n]; ok {
			mp[n] += 1
		} else {
			mp[n] = 1
		}
	}

	numberScore := make(numbers, len(mp)+1)
	for k, v := range mp {
		numberScore = append(numberScore, number{val: k, score: v * k})
	}
	sort.Sort(numberScore)

	dp := make([]int, len(numberScore))
	score := 0
	for i := 0; i < len(numberScore); i++ {
		if i < 2 {
			dp[i] = numberScore[i].score
		} else {
			needDelete := numberScore[i].val - 1
			_, ok := mp[needDelete]
			if ok {
				// 如果needDelete存在，那在numberScore数组中肯定和 i 相邻，并在前一个
				dp[i] = MaxInt(dp[i-1], dp[i-2]+numberScore[i].score)
			} else {
				// 不存在冲突相邻的，那就前两个都能取
				dp[i] = MaxInt(dp[i-1]+numberScore[i].score, dp[i-2]+numberScore[i].score)
			}

			if dp[i] > score {
				score = dp[i]
			}
		}
	}
	return score
}

const N = 10001

var (
	cnt [N]int
	f   [N][2]int
)

func deleteAndEarn_double100(nums []int) int {
	for i, _ := range cnt {
		cnt[i] = 0
	}

	for _, v := range nums {
		cnt[v]++
	}

	res := 0
	for i := 1; i < N; i++ {
		f[i][0] = MaxInt(f[i-1][0], f[i-1][1])
		f[i][1] = f[i-1][0] + i*cnt[i]
		res = MaxInt(f[i][0], f[i][1])
	}
	return res
}
