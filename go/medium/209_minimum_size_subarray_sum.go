func minSubArrayLen(target int, nums []int) int {
    minSize := math.MaxInt
    sum := nums[0]
    l, r := 0, 0
    for l < len(nums) && r < len(nums) {
        
        if sum < target {
            r++
            if r < len(nums) {
                sum += nums[r]
            }
            continue
        } 
        
        if sum >= target {
            minSize = min(minSize, r - l + 1)
            fmt.Println(l, r, minSize, sum)
            sum -= nums[l]
            l++
            continue
        }
    }

    if minSize == math.MaxInt {
        return 0
    }
    return minSize
}
