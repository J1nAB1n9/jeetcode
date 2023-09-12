package example1462

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	queriesSize := len(queries)
	answer := make([]bool, queriesSize)
	suffix := make(map[int][]int)
	var pd [][]bool
	for i := 0; i < numCourses; i++ {
		pd = append(pd, make([]bool, numCourses))
	}

	for _, v := range prerequisites {
		pd[v[0]][v[1]] = true
		suffix[v[0]] = append(suffix[v[0]], v[1])
	}

	for i := 0; i < numCourses; i++ {
		for j := 0; j < len(suffix[i]); j++ {
			v := suffix[i][j]
			for k := 0; k < numCourses; k++ {
				if pd[v][k] && !pd[i][k] {
					pd[i][k] = true
					suffix[i] = append(suffix[i], k)
				}
			}
		}
	}

	for i, q := range queries {
		answer[i] = pd[q[0]][q[1]]
	}

	return answer
}
