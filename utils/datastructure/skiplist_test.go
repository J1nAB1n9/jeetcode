package datastructure

import "testing"

func TestSkipListInteger_Insert(t *testing.T) {
	sl := NewSkipListInteger(5)

	sl.Insert(1, 1)
	sl.Insert(2, 1)
	sl.Insert(5, 1)
	sl.Insert(5, 2)
	sl.Insert(4, 1)
	sl.Insert(9, 1)
	sl.Insert(8, 1)
	sl.PrintfInfo()
}
