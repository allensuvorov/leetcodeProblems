func moveZeroes(nums []int)  {
    for plant, scout := 0, 0; scout < len(nums); scout++ {
        if nums[scout] != 0 {
            nums[plant], nums[scout] = nums[scout], nums[plant]
            plant++
        }
    }
}
