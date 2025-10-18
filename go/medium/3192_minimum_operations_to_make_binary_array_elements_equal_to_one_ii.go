func minOperations(nums []int) int {
    // - count alternations
    altCount := 0
    prev := 1
    for _, cur := range nums {
        if cur != prev {
            altCount++
            prev = cur
        }
    }
    return altCount
}
