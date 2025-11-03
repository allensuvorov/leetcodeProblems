func longestOnes(nums []int, k int) int {
    zeroCount := 0
    maxWindow := 0
    for l, r := 0, 0; r < len(nums); r++ {
        if nums[r] == 0 {
            zeroCount++
        }

        // cut tail if too many zeros
        for zeroCount > k {
            if nums[l] == 0 {
                zeroCount--
            }
            l++
        }

        currentWindow := r - l + 1
        maxWindow = max(maxWindow, currentWindow)
    }
    return maxWindow
}
