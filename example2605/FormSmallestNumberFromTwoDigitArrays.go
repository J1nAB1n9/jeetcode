package example2605
func minNumber(nums1 []int, nums2 []int) int {
	v1min := nums1[0]
	v2min := nums2[0]
	v12equip := 10
	//for _,v := range nums1 {
	//	if v1min > v {
	//		v1min = v
	//	}
	//}
	//
	//for _,v := range nums2 {
	//	if v2min > v {
	//		v2min = v
	//	}
	//}

	for i,v := range nums1 {
		if v1min > v {
			v1min = v
		}
		for _,w := range nums2 {
			if i == 0 && v2min > w {
				v2min = w
			}

			if v == w && v < v12equip {
				v12equip = w
			}
		}
	}

	if v12equip != 10 {
		return v12equip
	}

	if v1min < v2min {
		return v1min * 10 + v2min
	}

	return v2min * 10 + v1min
}