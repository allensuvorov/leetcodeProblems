func findMaxAverage(nums []int, k int) float64 {
    max := math.MinInt
    sum := 0

    for l, r := 0, 0; r < len(nums); r ++ {
        sum += nums[r]
        if r >= k - 1 {
            if sum > max {
                max = sum
            }
            sum -= nums[l]
            l++
        } 
    }
    return float64(max)/float64(k)
}
