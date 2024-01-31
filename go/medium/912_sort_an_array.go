func sortArray(nums []int) []int {
    // insertion sort
    for i := 1; i < len(nums); i++ {
        cur := nums[i]
        j := i
        for j > 0 && cur < nums[j - 1] {
            j--
            nums[j + 1] = nums[j]
        }
        nums[j] = cur
    }
    return nums
}
