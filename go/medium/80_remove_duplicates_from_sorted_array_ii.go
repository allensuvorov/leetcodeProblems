func removeDuplicates(nums []int) int {
    plant := 1
    occurance := 1

    for get := 1; get < len(nums); get++ {
        if nums[get] == nums[get - 1] {
            occurance++
        } else {
            occurance = 1
        }

        if occurance <= 2 {
            nums[plant] = nums[get]
            plant++
        }
    }
    return plant
}
