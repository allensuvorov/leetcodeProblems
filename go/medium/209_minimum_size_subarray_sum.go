func minSubArrayLen(target int, nums []int) int {
    sum := 0
    minSize := math.MaxInt

    for l, r := 0, 0; r < len(nums); r++ {
        sum += nums[r]

        for sum >= target {
            minSize = min(minSize, r - l + 1)
            sum -= nums[l]
            l++
        }
    }

    if minSize == math.MaxInt {
        return 0
    }
    return minSize
}
