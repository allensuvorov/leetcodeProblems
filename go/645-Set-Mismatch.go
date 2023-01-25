func findErrorNums(nums []int) []int {
    dup := -1
    missing := 1
    absV := 0
    for _, v := range nums {
        if v < 0 {
            absV = v * -1
        } else {
            absV = v
        }
        if nums[absV - 1] < 0 {
            dup = absV
        } else {
            nums[absV-1] *= -1
        }
    }
    for i, v := range nums{
        if v > 0 {
            missing = i + 1
        }
    }
    return []int{dup, missing}
}
