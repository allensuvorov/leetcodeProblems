func longestSubarray(nums []int) int {
    longestWindow := 0
    zeroCount := 0
    
    for tail, head := 0, 0; head < len(nums); head++ {
        if nums[head] == 0 {
            zeroCount++
        }

        for zeroCount > 1 {
            if nums[tail] == 0 {
                zeroCount--
            }
            tail++
        }
        longestWindow = max(longestWindow, head - tail)
    }
    return longestWindow
}
