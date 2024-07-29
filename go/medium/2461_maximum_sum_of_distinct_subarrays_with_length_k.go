func maximumSubarraySum(nums []int, k int) int64 {
    //  sliding window
    //  frequency map of repeating values
    //  set of distinct

    distinct := map[int]bool{}
    repeated := map[int]int{}
    maxSum := 0
    curSum := 0
    l := 0
    for r, v := range nums{
        if distinct[v] {
            distinct[v] = false
            repeated[v] = 2
        } else if repeated[v] > 0 {
            repeated[v]++
        } else {
            distinct[v] = true
        }
        curSum += v
        
        if r + 1 >= k {
            lv := nums[l]

            if len(repeated) == 0 {
                maxSum = max(maxSum, curSum)
                distinct[lv] = false
            } else {
                if distinct[lv] {
                    distinct[lv] = false
                } else if repeated[lv] > 2 {
                    repeated[lv]--
                } else if repeated[lv] == 2 {
                    delete(repeated, lv)
                    distinct[lv] = true
                }
            }
            curSum -= lv
            l++
        }
    }
    return int64(maxSum)
}
