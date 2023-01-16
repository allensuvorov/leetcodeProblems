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

// someone added a cleaner one
func containsNearbyDuplicate(nums []int, k int) bool {
    recorder := make(map[int]int)
    for i, num := range nums {
        j, ok := recorder[num]
        if ok && i - j <= k {
            return true
        }
        recorder[num] = i
    }
    return false
}
