func findDifference(nums1 []int, nums2 []int) [][]int {
    set1 := make(map[int]bool)
    set2 := make(map[int]bool)
    for _, v := range nums1 {
        set1[v] = true
    }
    for _, v := range nums2 {
        set2[v] = true
    }

    diff1 := []int{}
    for k := range set1 {
        if !set2[k] {
            diff1 = append(diff1, k)
        }
    }
    
    diff2 := []int{}
    for k := range set2 {
        if !set1[k] {
            diff2 = append(diff2, k)
        }
    }
    return [][]int{diff1, diff2}
}
