func minimumPairRemoval(nums []int) int {
// simulation
    opsCnt := 0
    for len(nums) > 1 {
        dec := false
        // check for dec
        for i := 1; i < len(nums); i++ {
            if nums[i-1] > nums[i] {
                dec = true
                break
            }
        }
        if !dec {
            return opsCnt
        }
        opsCnt++

        // find min sum pair
        minSum := nums[1] + nums[0]
        minSumIdx := 1
        for i := 1; i < len(nums); i++ {
            if nums[i] + nums[i-1] < minSum {
                minSum = nums[i] + nums[i-1]
                minSumIdx = i
            }
        }
        // sum the min sum pair, remove extra item
        nums[minSumIdx-1] += nums[minSumIdx]
        nums = append(nums[:minSumIdx], nums[minSumIdx + 1:]...)
    }
    return opsCnt
}
