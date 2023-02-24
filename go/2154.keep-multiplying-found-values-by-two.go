func findFinalValue(nums []int, original int) int {
    m := make(map[int]bool)
    for i := range nums {
        m[nums[i]] = true
    }
    n := original
    for {
        if m[n] {
            n *=2
        } else {
            return n
        }
    }
    return n
}
