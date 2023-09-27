package example1333

import "sort"

//restaurants[i] = [idi, ratingi, veganFriendlyi, pricei, distancei]
type restaurant struct {
	id   int
	rank int
}

type Restaurants []restaurant

func (p Restaurants) Len() int      { return len(p) }
func (p Restaurants) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Restaurants) Less(i, j int) bool {
	if p[i].rank == p[j].rank {
		return p[i].id > p[j].id
	}
	return p[i].rank > p[j].rank
}

func (p Restaurants) GetInes() []int {
	ans := make([]int, p.Len())
	for i, v := range p {
		ans[i] = v.id
	}

	return ans
}

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var rs Restaurants
	for _, v := range restaurants {
		if v[2] == 0 && veganFriendly == 1 {
			continue
		}
		if v[3] > maxPrice || v[4] > maxDistance {
			continue
		}
		r := restaurant{
			id:   v[0],
			rank: v[1],
		}
		rs = append(rs, r)
	}
	sort.Sort(rs)

	return rs.GetInes()
}
