package example2258

const (
	Grass = int(0)
	Flame = int(1)
	Fence = int(2)

	noAns = 1000000000
)

type index struct {
	i, j int
	ans  int
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
	safeTime := make(map[int]int)

	var flameQueue []index
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == Flame {
				fire := index{i: i, j: j}
				safeTime[fire.i*1000+fire.j] = 0
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

			if _, ok := safeTime[nextTime.i*1000+nextTime.j]; ok {
				continue
			}

			safeTime[nextTime.i*1000+nextTime.j] = safeTime[flameIndex.i*1000+flameIndex.j] + 1
			flameQueue = append(flameQueue, nextTime)
		}
	}

	flameQueue = flameQueue[:0]
	begin := index{0, 0, 0xfffffffe}
	flameQueue = append(flameQueue, begin)

	safeLeave := false
	var ansNode index
	// ans := 0xfffffffe
	hereTime := make(map[int]int)
	hereTime[begin.i*1000+begin.j] = 0
	for len(flameQueue) != 0 {
		nowPos := flameQueue[0]
		flameQueue = flameQueue[1:]
		for k := 0; k < 4; k++ {
			x := nowPos.i + nl[k]
			y := nowPos.j + ml[k]
			nextPos := index{
				i:   x,
				j:   y,
				ans: nowPos.ans,
			}
			if !checkGrid(x, y, n, m) {
				continue
			}

			if grid[x][y] != Grass {
				continue
			}

			// 已经维护过
			if _, ok := hereTime[nextPos.i*1000+nextPos.j]; ok {
				if x != n-1 || y != m-1 {
					continue
				}
			}

			nextTime := hereTime[nowPos.i*1000+nowPos.j] + 1
			if fireTime, has := safeTime[nextPos.i*1000+nextPos.j]; has {
				if nextTime > fireTime {
					continue
				} else if nextTime == fireTime && (x != n-1 || y != m-1) {
					continue
				}

				if (fireTime - nextTime) < nextPos.ans {
					nextPos.ans = fireTime - nextTime
					if x == n-1 && y == m-1 {
						nextPos.ans = nextPos.ans + 1
					}
				}
			}

			hereTime[nextPos.i*1000+nextPos.j] = nextTime

			if x == n-1 && y == m-1 {
				if safeLeave == false {
					safeLeave = true
					ansNode = nextPos
				} else if ansNode.ans < nextPos.ans {
					ansNode = nextPos
				}

				// break
			} else {
				flameQueue = append(flameQueue, nextPos)
			}
		}
		//if safeLeave {
		//	break
		//}
	}

	if !safeLeave {
		return -1
	} else if ansNode.ans == 0xfffffffe {
		return noAns
	}

	return ansNode.ans - 1
}

func MainFunc(g [][]int) int {
	return maximumMinutes(g)
}
