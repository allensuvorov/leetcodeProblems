func countElements(nums []int) int {
    // exclude min and max
    min := math.MaxInt
    max := math.MinInt
    for _, v := range nums {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }
    count := 0
    for _, v := range nums {
        if v > min && v < max {
            count++
        }
    }
    return count
}
