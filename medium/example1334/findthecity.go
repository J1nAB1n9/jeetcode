package example1334

type dis struct {
	index int
	dis   int
}
type node struct {
	index int
	next  []dis
}

var tree []node
var isOld map[int]struct{}
var getOld map[int]struct{}

func find(n node, now, distanceThreshold int, max int) {
	for _, v := range n.next {
		if max < len(getOld) {
			return
		}

		if v.dis+now > distanceThreshold {
			continue
		}
		if _, ok := isOld[v.index]; ok {
			continue
		}

		getOld[v.index] = struct{}{}
		isOld[v.index] = struct{}{}

		find(tree[v.index], now+v.dis, distanceThreshold, max)
		delete(isOld, v.index)
	}
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	tree = make([]node, n)
	for _, v := range edges {
		l := v[0]
		r := v[1]
		d := v[2]

		tree[l].next = append(tree[l].next, dis{
			index: r,
			dis:   d,
		})
		tree[r].next = append(tree[r].next, dis{
			index: l,
			dis:   d,
		})
	}

	ansIndex := -1
	ansNodeNum := 0xffffffff
	for i := n - 1; i >= 0; i-- {
		isOld = make(map[int]struct{})
		isOld[i] = struct{}{}

		getOld = make(map[int]struct{})
		getOld[i] = struct{}{}

		find(tree[i], 0, distanceThreshold, ansNodeNum)
		if ansNodeNum >= len(getOld) {
			ansNodeNum = len(getOld)
			ansIndex = i
		}

		// fmt.Printf("index:%d node:%d\n", i, len(getOld))

	}
	return ansIndex
}
