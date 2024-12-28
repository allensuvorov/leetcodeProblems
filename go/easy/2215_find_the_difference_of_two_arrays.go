func findDifference(nums1 []int, nums2 []int) [][]int {
    set1 := map[int]bool{}
    for _, v := range nums1 {
        set1[v] = true
    }

    set2 := map[int]bool{}
    for _, v := range nums2 {
        set2[v] = true
    }

    res := make([][]int, 2)
    
    for v := range set1 {
        if !set2[v] {
            res[0] = append(res[0], v)
        }
    }

    for v := range set2 {
        if !set1[v] {
            res[1] = append(res[1], v)
        }
    }

    return res
}
