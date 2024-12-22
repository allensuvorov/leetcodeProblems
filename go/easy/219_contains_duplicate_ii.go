func containsNearbyDuplicate(nums []int, k int) bool {
    numIdx := map[int]int{}
    for i, v := range nums {
        if j, ok := numIdx[v]; ok && i - j <= k {
            return true
        }
        numIdx[v] = i
    }
    return false
}
