func getCommon(nums1 []int, nums2 []int) int {
    i1, i2 := 0, 0
    for i1 < len(nums1) && i2 < len(nums2) {
        if nums1[i1] == nums2[i2] {
            return nums1[i1]
        }
        if nums1[i1] > nums2[i2] {
            i2++
        } else {
            i1++
        }
    }
    return -1
}
