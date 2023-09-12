package example1462

import "container/list"

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	queriesSize := len(queries)
	answer := make([]bool, queriesSize)

	var pd [][]bool
	for i := 0; i < numCourses; i++ {
		pd = append(pd, make([]bool, numCourses))
	}

	for _, v := range prerequisites {
		pd[v[0]][v[1]] = true
	}

	for i, q := range queries {
		_q0 := q[0]
		_q1 := q[1]

		if pd[_q0][_q1] {
			answer[i] = true
			continue
		}

		list.New()
	}

	return answer
}
