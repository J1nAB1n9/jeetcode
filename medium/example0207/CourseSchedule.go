package example0207

type node struct {
	val int

	in  int
	out []int
}

// 拓扑排序
// 时间 12ms 	击败 56.97%
// 内存 5.91MB 	击败 57.19%
func canFinish(numCourses int, prerequisites [][]int) bool {
	mm := make([]*node, numCourses)
	if len(prerequisites) == 0 {
		return true
	}

	sizeTuoppu := 0
	for _, v := range prerequisites {
		inNode := v[0]
		ouNode := v[1]

		if mm[inNode] == nil {
			mm[inNode] = &node{
				val: inNode,
				in:  0,
				out: nil,
			}
			sizeTuoppu += 1
		}

		if mm[ouNode] == nil {
			mm[ouNode] = &node{
				val: ouNode,
				in:  0,
				out: nil,
			}
			sizeTuoppu += 1
		}

		mm[inNode].out = append(mm[inNode].out, ouNode)
		mm[ouNode].in += 1
	}

	for i := 0; i < sizeTuoppu; i++ {
		find := false
		for j := 0; j < numCourses; j++ {
			if mm[j] == nil {
				continue
			}

			if mm[j].in == 0 {
				find = true
				for _, n := range mm[j].out {
					mm[n].in -= 1
				}
				// fmt.Println(mm[j])
				mm[j] = nil
				break
			}
		}

		if !find {
			return false
		}
	}

	return true
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	mm := make([]*node, numCourses)
	if len(prerequisites) == 0 {
		return true
	}

	sizeTuoppu := 0
	for _, v := range prerequisites {
		inNode := v[0]
		ouNode := v[1]

		if mm[inNode] == nil {
			mm[inNode] = &node{
				val: inNode,
				in:  0,
				out: nil,
			}
			sizeTuoppu += 1
		}

		if mm[ouNode] == nil {
			mm[ouNode] = &node{
				val: ouNode,
				in:  0,
				out: nil,
			}
			sizeTuoppu += 1
		}

		mm[inNode].out = append(mm[inNode].out, ouNode)
		mm[ouNode].in += 1
	}

	for i := 0; i < sizeTuoppu; i++ {
		find := false
		for j := 0; j < numCourses; j++ {
			if mm[j] == nil {
				continue
			}

			if mm[j].in == 0 {
				find = true
				for _, n := range mm[j].out {
					mm[n].in -= 1
				}
				// fmt.Println(mm[j])
				mm[j] = nil
				break
			}
		}

		if !find {
			return false
		}
	}

	return true
}
