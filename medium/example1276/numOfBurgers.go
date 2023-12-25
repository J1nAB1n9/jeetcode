package example1276

// 4X + 2Y = tomatoSlices
//  X +  Y = cheeseSlices
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	x := (tomatoSlices - cheeseSlices*2) / 2
	y := cheeseSlices - x

	if x+y != cheeseSlices || (4*x+2*y) != tomatoSlices || x < 0 || y < 0 {
		return []int{}
	}

	return []int{x, y}
}
