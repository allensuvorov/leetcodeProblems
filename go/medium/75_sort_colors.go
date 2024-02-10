func sortColors(nums []int)  {
    pivot := 1
    l, m, r := 0, 0, len(nums) - 1
    for m <= r {
        switch {
        case nums[m] < pivot:
            nums[l], nums[m] = nums[m], nums[l]
            l++
            m++
        case nums[m] > pivot:
            nums[r], nums[m] = nums[m], nums[r]
            r--
        default:
            m++
        }
    }    
}
