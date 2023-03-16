func targetIndices(nums []int, target int) []int {
    lessCount := 0
    targetCount := 0
    for i := range nums {
        if nums[i] == target {
            targetCount ++
        }
        if nums[i] < target {
            lessCount ++
        }
    }
    result := make([]int, targetCount)
    for i := range result {
        result[i] = lessCount + i 
    }
    return result
}
