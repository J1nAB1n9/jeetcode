package example1129

type node struct {
	red  []int
	blue []int
}

type flag struct {
	red  bool
	blue bool
}

var tree []*node
var paths []int
var findMap map[int]flag

func getPaths(red bool, long int, index int) {
	exit := findMap[index]
	if red {
		if exit.red {
			return
		}
		exit.red = true
	} else {
		if exit.blue {
			return
		}
		exit.blue = true
	}

	if paths[index] == -1 {
		paths[index] = long
	} else if paths[index] > long {
		paths[index] = long
	} else {
		// pass
	}

	if red {
		for _, v := range tree[index].red {
			getPaths(!red, long+1, v)
		}
	} else {
		for _, v := range tree[index].blue {
			getPaths(red, long+1, v)
		}
	}

}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	tree = make([]*node, n)
	paths = make([]int, n)
	findMap = make(map[int]flag)

	// init paths and findMap
	for i := 0; i < n; i++ {
		if i == 0 {
			paths[i] = 0
			findMap[i] = flag{
				red:  true,
				blue: true,
			}
		} else {
			paths[i] = -1
			findMap[i] = flag{
				red:  false,
				blue: false,
			}
		}

		nn := &node{}
		tree = append(tree, nn)
	}

	for i := 0; i < len(redEdges); i++ {
		// tree[redEdges[i][0]].red = append(tree[redEdges[i][0]].red, redEdges[i][1])
	}
	for i := 0; i < len(blueEdges); i++ {
		// tree[blueEdges[i][0]].red = append(tree[blueEdges[i][0]].red, blueEdges[i][1])
	}

	for _, v := range tree[0].red {
		getPaths(false, 1, v)
	}

	for _, v := range tree[0].blue {
		getPaths(true, 1, v)
	}

	return paths
}
