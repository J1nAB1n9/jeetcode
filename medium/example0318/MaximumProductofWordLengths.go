package example0318

import "sort"

type word struct {
	val  int32
	size int32
}
type ws []word

func (w ws) Len() int {
	return len(w)
}
func (w ws) Less(i, j int) bool {
	return w[i].size < w[j].size
}

func (w ws) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func maxProduct(words []string) int {
	arr := make(ws, len(words))
	isBad := int32(0)
	for i, v := range words {
		w := word{
			val:  0,
			size: 0,
		}
		for _, c := range v {
			index := 1 << int(c-'a')
			w.val |= int32(index)
			w.size += 1
		}

		arr[i] = w
		if isBad == 0 {
			isBad = w.val
		} else {
			isBad = isBad & isBad
		}
	}

	sort.Sort(arr)

	ans := 0
	for i := 0; i < len(arr); i++ {
		// 如果当前平方已经小于最大值，那就可以剪了
		if int(arr[i].size*arr[i].size) < ans {
			break
		}

		for j := i + 1; j < len(arr); j++ {
			// 剪枝
			if int(arr[i].size*arr[j].size) < ans {
				break
			}

			if arr[i].val&arr[j].val != 0 {
				continue
			}

			ans = int(arr[i].size * arr[j].size)
			break
		}
	}
	return ans
}
