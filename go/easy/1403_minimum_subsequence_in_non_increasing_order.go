func minSubsequence(nums []int) []int {
    res := []int{}
    // get sum
    sum := 0
    for _, v := range nums {
        sum += v
    }
    // sort
    sort.Ints(nums)
    subSum := 0
    for i := len(nums) - 1; 2* subSum <= sum ; i-- {
        subSum += nums[i]
        res = append(res, nums[i])
    }
    return res
}
