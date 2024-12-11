func removeElement(nums []int, val int) int {
    plant := 0
    for scout := 0; scout < len(nums): scout++ {
        if nums[scout] != val {
            nums[plant] = nums[scout]
            plant++
        }
    }
    return plant
}
