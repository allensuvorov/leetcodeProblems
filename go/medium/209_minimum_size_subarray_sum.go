func minSubArrayLen(target int, nums []int) int {
    ans := len(nums) + 1
    l, r := 0, 0
    sum := 0
    for r < len(nums) {
        sum += nums[r]
        for l < len(nums) && sum >= target {
            if r - l + 1 < ans {
                ans = r - l + 1
            }
            sum -= nums[l]
            l++
        }
        r++
    }
    if ans == len(nums) + 1 {
        return 0
    }
    return ans
}
