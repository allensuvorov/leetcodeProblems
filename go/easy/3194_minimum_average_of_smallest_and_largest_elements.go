func minimumAverage(nums []int) float64 {
    minAve := float64(50)
    slices.Sort(nums)
    for l, r := 0, len(nums) - 1; l < r; {
        minAve = min(minAve, (float64(nums[l]) + float64(nums[r])) / 2)
        l++
        r--
    }
    return minAve
}
