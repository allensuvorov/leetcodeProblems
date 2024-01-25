func sortColors(nums []int)  {
    r := len(nums) - 1
    for r > 0 {
        for l := 0; l < r; l++ {
            if nums[l] > nums[l + 1] {
                nums[l], nums[l + 1] = nums[l + 1], nums[l]
            }
        }
        r--
    }
}
