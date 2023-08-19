func containsNearbyDuplicate(nums []int, k int) bool {
    m := map[int]int{}
    for i, v := range nums {
        if j, ok := m[v]; ok {
            if i - j <= k {
                return true
            } 
        }
        m[v] = i
    }
    return false
}
