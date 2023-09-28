package example2251

import "sort"

//时间 172ms
//击败 100.00%使用 Go 的用户
//内存 12.87MB
//击败 85.71%使用 Go 的用户

type flower struct {
	tm int
	be int
}
type FlowerArr []flower

func (a FlowerArr) Len() int           { return len(a) }
func (a FlowerArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FlowerArr) Less(i, j int) bool { return a[i].tm < a[j].tm }

type human struct {
	tm  int
	inx int
}

type humanArr []human

func (a humanArr) Len() int           { return len(a) }
func (a humanArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a humanArr) Less(i, j int) bool { return a[i].tm < a[j].tm }

func fullBloomFlowers(flowers [][]int, people []int) []int {
	fs := make([]flower, 0, len(flowers)*2+1)
	// init arr
	for _, v := range flowers {
		if v[0] == v[1] {
			fs = append(fs, flower{
				tm: v[0],
				be: 2,
			})
			continue
		}

		fs = append(fs, flower{
			tm: v[0],
			be: 1,
		})
		fs = append(fs, flower{
			tm: v[1],
			be: 0,
		})
	}
	humans := make([]human, len(people))
	for i, v := range people {
		humans[i].tm = v
		humans[i].inx = i
	}
	// keep sort
	sort.Sort(FlowerArr(fs))
	sort.Sort(humanArr(humans))

	flowerNum := 0
	fIndex := 0

	lastEquipVal := -1
	lastEquipInx := -1

	for _, v := range humans {
		lastFlower := 0
		for fIndex < len(fs) && fs[fIndex].tm < v.tm {
			if fs[fIndex].be == 1 {
				flowerNum += 1
			} else if fs[fIndex].be == 0 {
				flowerNum -= 1
			}
			fIndex++
		}

		if lastEquipInx == v.tm {
			lastFlower = lastEquipVal
		} else {
			ise := false
			for tmp := fIndex; tmp < len(fs) && fs[tmp].tm == v.tm; tmp++ {
				ise = true
				if fs[tmp].be == 0 {
					continue
				}

				lastFlower += 1
			}

			if ise == true {
				lastEquipVal = lastFlower
				lastEquipInx = v.tm
			}
		}

		people[v.inx] = flowerNum + lastFlower
	}

	return people
}

func FullBloomFlowers(flowers [][]int, people []int) []int {
	return fullBloomFlowers(flowers, people)
}
