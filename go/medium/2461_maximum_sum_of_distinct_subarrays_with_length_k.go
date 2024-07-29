func maximumSubarraySum(nums []int, k int) int64 {
    numsFreqs := [10e5+1]int{}
    var repeatedCount, maxSum, curSum int
    for r, rv := range nums{
        if numsFreqs[rv] == 1 {
            repeatedCount++
        }
        numsFreqs[rv]++
        curSum += rv
        if l := r - k + 1; l >= 0 {
            lv := nums[l]
            if repeatedCount == 0 {
                maxSum = max(maxSum, curSum)
            } 
            if numsFreqs[lv] == 2 {
                repeatedCount--
            }
            numsFreqs[lv]--
            curSum -= lv
        }
    }
    return int64(maxSum)
}
