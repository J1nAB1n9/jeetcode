package example3208

var l int
var r int

var size int

func _get() int {
	if r-1 < 0 {
		return size - 1
	}
	return r - 1
}

func _getlen() int {
	if l > r {
		return r + size - l + 1
	}

	return r - l + 1
}

func numberOfAlternatingGroups(colors []int, k int) int {
	l = 0
	size = len(colors)
	r = l + 1%size

	ans := 0
	for l < size {
		if colors[r] != colors[_get()] {
			if _getlen() == k {
				ans++
				l++
				r = (r + 1) % size
			} else {
				r = (r + 1) % size
			}
		} else {
			if l == _get() {
				l++
				r = (r + 1) % size
			} else {
				l += 1
			}
		}
	}
	return ans
}
