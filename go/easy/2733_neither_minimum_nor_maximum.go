func findNonMinOrMax(nums []int) int {
    max, min := 1, 100
    for _, v := range nums {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }
    for _, v := range nums {
        if v != max && v != min {
            return v
        }
    }
    return -1
}
