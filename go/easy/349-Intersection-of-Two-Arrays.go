func intersection(nums1 []int, nums2 []int) []int {
    m := make(map[int]bool)
    for _, v := range nums1 {
        m[v]=true
    }
    
    nums3 := []int{}
    for _, v := range nums2 {
        if m[v] {
            nums3 = append (nums3, v)
            m[v] = false
        }
    }
    return nums3
}
