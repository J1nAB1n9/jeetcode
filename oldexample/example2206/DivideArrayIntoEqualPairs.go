package example2206

func divideArray(nums []int) bool {
    mm := make(map[int]int)

    for _,v := range nums {
        mm[v]++
    }

    for _,v := range mm {
        if v % 2 == 1 {
            return false
        }
    }

    return true
}