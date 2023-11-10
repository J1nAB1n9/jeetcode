package example2258

const (
	Grass = int(0)
	Flame = int(1)
	Fence = int(2)

	noAns = 1000000000
)

type index struct {
	i, j int
}

var nl = []int{1, 0, -1, 0}
var ml = []int{0, -1, 0, 1}

func checkGrid(i, j, n, m int) bool {
	if i < 0 || i >= n {
		return false
	}

	if j < 0 || j >= m {
		return false
	}

	return true
}

func maximumMinutes(grid [][]int) int {
	if grid[0][0] != Grass {
		return noAns
	}

	n, m := len(grid), len(grid[0])
	safeTime := make(map[index]int)

	var flameQueue []index
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == Flame {
				fire := index{i: i, j: j}
				safeTime[fire] = 0
				flameQueue = append(flameQueue, fire)
			}
		}
	}

	for len(flameQueue) != 0 {
		flameIndex := flameQueue[0]
		flameQueue = flameQueue[1:]

		for k := 0; k < 4; k++ {
			x := flameIndex.i + nl[k]
			y := flameIndex.j + ml[k]
			if !checkGrid(x, y, n, m) {
				continue
			}

			if grid[x][y] != Grass {
				continue
			}

			nextTime := index{
				i: x,
				j: y,
			}

			if _, ok := safeTime[nextTime]; ok {
				continue
			}

			safeTime[nextTime] = safeTime[flameIndex] + 1
			flameQueue = append(flameQueue, nextTime)
		}
	}

	flameQueue = flameQueue[:0]
	begin := index{0, 0}
	flameQueue = append(flameQueue, begin)

	safeLeave := false
	ans := 0xfffffffe
	hereTime := make(map[index]int)
	hereTime[begin] = 0
	for len(flameQueue) != 0 {
		nowPos := flameQueue[0]
		flameQueue = flameQueue[1:]
		for k := 0; k < 4; k++ {
			x := nowPos.i + nl[k]
			y := nowPos.j + ml[k]
			nextPos := index{
				i: x,
				j: y,
			}
			if !checkGrid(x, y, n, m) {
				continue
			}

			if grid[x][y] != Grass {
				continue
			}

			// 已经维护过
			if _, ok := hereTime[nextPos]; ok {
				continue
			}

			nextTime := hereTime[nowPos] + 1
			if fireTime, has := safeTime[nextPos]; has {
				if nextTime >= fireTime {
					continue
				}

				if (fireTime - nextTime) < ans {
					ans = fireTime - nextTime
				}
			}

			hereTime[nextPos] = nextTime
			flameQueue = append(flameQueue, nextPos)
			if x == n-1 && y == m-1 {
				safeLeave = true
				break
			}
		}
		if safeLeave {
			break
		}
	}

	if !safeLeave {
		return -1
	} else if ans == 0xfffffffe {
		return noAns
	}
	return ans
}

func MainFunc(g [][]int) int {
	return maximumMinutes(g)
}
