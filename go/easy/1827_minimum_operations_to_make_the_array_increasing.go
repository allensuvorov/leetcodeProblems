func minOperations(nums []int) int {
    ans := 0
    if len(nums) == 1 {
        return 0
    }
    for i := 1; i < len(nums); i++ {
        diff := nums[i] - nums[i-1]
        switch {
        case diff == 0:
            ans++
            nums[i]++
        case diff < 0:
            ans += -1*diff + 1
            nums[i] += -1*diff + 1
        }
    }
    return ans
}
