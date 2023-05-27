func minStartValue(nums []int) int {
    baseline := 1
    sum := baseline
    min := 1
    for i := range nums {
        sum += nums[i]
        if min > sum {
            min = sum
        }
    }
    switch {
    case min >= baseline:
        return baseline
    case min < baseline:
        return baseline + (baseline - min)
    }
    return 0
}
