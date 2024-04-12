func isPossibleToSplit(nums []int) bool {
    fm := make([]int, 101)
    for _, v := range nums {
        fm[v]++
        if fm[v] > 2 {
            return false
        }
    }
    return true
}
