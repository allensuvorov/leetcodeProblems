func containsNearbyDuplicate(nums []int, k int) bool {
    hm := map[int]int{}
    for i, v := range nums {
        if _, ok := hm[v]; !ok {
            hm[v] = i
            continue
        }
        if i - hm[v] <= k {
            return true
        } else {
            hm[v] = i
        }
    }
    return false
}
