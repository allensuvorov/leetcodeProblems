func longestOnes(nums []int, k int) int {
    zeroCount := 0
    maxLen := 0
    for l, r := 0, 0; r < len(nums); r++ {
        if nums[r] == 0 {
            zeroCount++
        }
        
        for zeroCount > k {
            if nums[l] == 0 {
                zeroCount--
            }
            l++
        }
        maxLen = max(maxLen, r - l + 1)
    }
    return maxLen
}
