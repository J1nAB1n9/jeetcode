package example0765

var f []int

func find(i int) int {
	if i != f[i] {
		f[i] = find(f[i])
	}

	return f[i]
}

func union(i, j int) {
	v, w := find(i), find(j)
	if v != w {
		f[v] = f[w]
	}
}

func minSwapsCouples(row []int) int {
	n := len(row)
	f = make([]int, n)
	for i := 0; i < n; i += 2 {
		f[i] = i
		f[i+1] = i
	}

	//for i := 0; i < n; i += 2 {
	//	union(i, i+1)
	//}
	count := 0
	for i := 0; i < n; i += 2 {
		if find(row[i]) != find(row[i+1]) {
			count += 1
		}
		// union(find(row[i]), find(row[i+1]))
	}
	if count == 0 {
		return 0
	}

	for i := 0; i < n; i += 2 {
		union(find(row[i]), find(row[i+1]))
	}
	mn := make(map[int]int)
	des := 0
	for i := 0; i < n; i++ {
		v := find(i)
		mn[v] += 1
		if mn[v] == 3 {
			des += 1
		}
	}

	return count - des
}
