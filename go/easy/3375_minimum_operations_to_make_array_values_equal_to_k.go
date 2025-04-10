func minOperations(nums []int, k int) int {
    set := make(map[int]bool)

    for _, v := range nums {
        if v < k {
            return -1
        }
        if v > k {
            set[v] = true
        }
    }

    return len(set)
}
