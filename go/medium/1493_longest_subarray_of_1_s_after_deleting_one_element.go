func longestSubarray(nums []int) int {
    maxSize := 0
    zeroCount := 0
    l := 0
    for r := range nums {
        if nums[r] == 0 {
            zeroCount++
        }
        for zeroCount > 1 {
            if nums[l] == 0 {
                zeroCount--
            }
            l++
        }
        maxSize = max(maxSize, r - l) // (r - l) is windowSize minus 1
            // when there is one a more zeros in the array, the maxSize window will include one of them
            // when there no zeros in the array, the maxSize be reduced by one, since we have to delete one
    }   
    return maxSize
}
