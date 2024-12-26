func moveZeroes(nums []int)  {
    for i := range nums {
        if nums[i] != 0 {
            continue
        }

        for j := i + 1; j < len(nums); j++  {
            if nums[j] != 0 {
                nums[i], nums[j] = nums[j], nums[i]
                break
            }
        }
    }
}
