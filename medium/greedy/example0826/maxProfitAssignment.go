package example0826

import "sort"

type work struct {
	difficulty int
	profit     int
}

type works []work

func (w works) Len() int           { return len(w) }
func (w works) Less(i, j int) bool { return w[i].difficulty < w[j].difficulty }
func (w works) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	ws := make(works, len(difficulty))

	for i := 0; i < len(difficulty); i++ {
		ws[i] = work{
			difficulty: difficulty[i],
			profit:     profit[i],
		}
	}

	sort.Sort(ws)
	// worker sort
	sort.Ints(worker)

	index := 0
	ans := 0
	maxProfit := 0
	for _, wr := range worker {
		for ; index < len(ws) && wr >= ws[index].difficulty; index++ {
			if maxProfit < ws[index].profit {
				maxProfit = ws[index].profit
			}
		}
		ans += maxProfit
	}
	return ans
}
