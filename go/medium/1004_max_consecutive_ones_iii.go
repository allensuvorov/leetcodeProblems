func longestOnes(nums []int, k int) int {
    maxOneCount := 0
    zeroCount := 0
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
        maxOneCount = max(maxOneCount, r - l + 1)
    }
    return maxOneCount
}
