func sortEvenOdd(nums []int) []int {
    odds := make([]int, 101)
    evens := make([]int, 101)
    for i, v := range nums {
        if i%2 == 0 {
            evens[v]++
        } else {
            odds[v]++
        }
    }
    l := 1
    r := 100
    for i := range nums {
        if i%2 == 0 {
            for evens[l] == 0 {
                l++
            }
            nums[i] = l
            evens[l]--
        } else {
            for odds[r] == 0 {
                r--
            }
            nums[i] = r
            odds[r]--
        }
    }
    return nums
}
