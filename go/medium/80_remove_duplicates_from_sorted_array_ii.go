func removeDuplicates(nums []int) int {
    plant := 1
    repCount := 0

    for get := 1; get < len(nums); get++ {
        if nums[get] == nums[get - 1] {
            repCount++
        } else {
            repCount = 0
        }

        if repCount < 2 {
            nums[plant] = nums[get]
            plant++
        }
    }
    return plant
}
