package example2605

import (
	"fmt"
	"testing"
)

//Input: nums1 = [4,1,3], nums2 = [5,7]
//Output: 15
//Explanation: The number 15 contains the digit 1 from nums1 and the digit 5 from nums2. It can be proven that 15 is the smallest number we can have.

//Input: nums1 = [3,5,2,6], nums2 = [3,1,7]
//Output: 3
//Explanation: The number 3 contains the digit 3 which exists in both arrays.

func TestExample(t *testing.T) {
	fmt.Println(minNumber([]int{4,1,3},[]int{5,7}))
	fmt.Println(minNumber([]int{3,5,2,6},[]int{3,1,7}))
}
