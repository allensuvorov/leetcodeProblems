func moveZeroes(nums []int)  {
    for insertIndex, scout := 0, 0; scout < len(nums); scout++ {
        if nums[scout] != 0 {
            nums[insertIndex], nums[scout] = nums[scout], nums[insertIndex]
            insertIndex++
        }
    }
}
