package example1669

import "testing"

//[0,1,2,3,4,5]
//3
//4
//[1000000,1000001,1000002]
func TestDemo1(t *testing.T) {
	var list1Num []int = []int{0, 1, 2, 3, 4, 5}
	var list2Num []int = []int{1000000, 1000001, 1000002}

	list1 := &ListNode{}
	list2 := &ListNode{}
	ans := MergeInBetween(AddValList(list1, list1Num), 3, 4, AddValList(list2, list2Num))
	ShowVal(ans, true)
}
