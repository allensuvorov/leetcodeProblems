func maximumTripletValue(nums []int) int64 {
    var maxVal int64
    var curVal int64
    for i := range nums {
        curVal += int64(nums[i])
        for j := i + 1; j < len(nums) - 1; j++ {
            curVal -= int64(nums[j])
            for k := j + 1; k < len(nums); k++ {
                curVal *= int64(nums[k])
                maxVal = max(maxVal, curVal)
                curVal /= int64(nums[k])
            }
            curVal += int64(nums[j])
        }
        curVal -= int64(nums[i])
    }
    return maxVal
}
