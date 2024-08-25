func largestSumAfterKNegations(nums []int, k int) int {
    slices.Sort(nums)
    
    for i := 0; i < len(nums) && nums[i] < 0 && k > 0; i++ {
        nums[i] = -nums[i]
        k--
    }

    slices.Sort(nums)

    if nums[0] > 0 && k > 0 && k % 2 == 1 {
        nums[0] = - nums[0]
    }

    sum := 0
    for _, v := range nums {
        sum += v
    }

    return sum
}
