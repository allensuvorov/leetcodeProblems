func findDifference(nums1 []int, nums2 []int) [][]int {
    set1 := make(map[int]bool)
    set2 := make(map[int]bool)

    for _, v := range nums1 {
        set1[v] = true
    }

    for _, v := range nums2 {
        set2[v] = true
    }
    
    // check what is missing in nums2 (set2)
    result1 := []int{}
    for v := range set1 {
        if !set2[v] {
            result1 = append(result1, v)
        }
    }

    // check what is missing in nums1 (set1)
    result2 := []int{}
    for v := range set2 {
        if !set1[v] {
            result2 = append(result2, v)
        }
    }
    return [][]int{result1, result2}
}
