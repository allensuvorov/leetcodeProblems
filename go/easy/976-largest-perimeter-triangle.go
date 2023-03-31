func largestPerimeter(nums []int) int {
    sort.Ints(nums)
    for i := len(nums)-1; i > 1; i -- {
        if nums[i]/2 < nums[i-1] {
            if nums[i-1] > nums[i] - nums [i-2] {
                return nums[i-1] + nums[i] + nums[i-2]
            }
        }
    }
    return 0
}
