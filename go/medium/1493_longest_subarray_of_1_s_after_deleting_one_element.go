func longestSubarray(nums []int) int {
    maxLen := 0
    zeroCount := 0
    for l, r := 0, 0; r < len(nums); r++ {
        if nums[r] == 0 {
            zeroCount++
        }

        for zeroCount > 1 {
            if nums[l] == 0 {
                zeroCount--
            }
            l++
        }
        maxLen = max(maxLen, r - l)
    }
    return maxLen   
}
