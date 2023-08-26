func minSubArrayLen(target int, nums []int) int {
    ans := len(nums)
    l, r := 0, 0
    sum := 0
    exists := false
    for r < len(nums) {
        sum += nums[r]
        for l < len(nums) && sum >= target {
            exists = true
            if r - l + 1 < ans {
                ans = r - l + 1
            }
            sum -= nums[l]
            l++
        }
        r++
    }
    if !exists {
        return 0
    }
    return ans
}
