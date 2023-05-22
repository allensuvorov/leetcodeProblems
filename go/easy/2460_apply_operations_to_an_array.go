func applyOperations(nums []int) []int {
    // do operations
    for i := 0; i < len(nums)-1; i ++ {
        if nums[i] == nums[i+1] {
            nums[i] *= 2
            nums[i+1] = 0
        }
    }
    // move zeros
    j := 0 
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == 0 {
            
            if j == 0 {
                j = i + 1
            }
            for j < len(nums) {
                if nums[j] != 0 {
                    nums[i] = nums[j]
                    nums[j] = 0
                    j++
                    break
                }
                j++
                if j == len(nums) {
                    return nums
                }
            }
        }
    }
    return nums
}
