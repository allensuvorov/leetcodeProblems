func merge(nums1 []int, m int, nums2 []int, n int)  {
    for i:= len(nums1)-1; i>=0; i--{
        if n == 0 {
            break
        }
        if m == 0 || nums1[m-1] < nums2[n-1] {
            nums1[i] = nums2[n-1]
            n--
            continue
        }
        nums1[i] = nums1[m-1]
        m--
    }
    return
}
