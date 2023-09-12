package example0210

type node struct {
	val int

	in  int
	out []int
}

func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
	mm := make([]*node, numCourses)
	//if len(prerequisites) == 0 {
	//	return
	//}

	sizeTuoppu := 0
	for _, v := range prerequisites {
		inNode := v[1]
		ouNode := v[0]

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
			if mm[j] == nil || mm[j].val == -1 {
				continue
			}

			if mm[j].in == 0 {
				find = true
				for _, n := range mm[j].out {
					mm[n].in -= 1
				}
				// fmt.Println(mm[j])
				ans = append(ans, mm[j].val)
				mm[j].val = -1
				break
			}
		}

		if !find {
			return []int{}
		}
	}

	if len(ans) != numCourses {
		for i, v := range mm {
			if v == nil {
				ans = append(ans, i)
			}
		}
	}

	return
}
