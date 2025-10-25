func increasingTriplet(nums []int) bool {
    min1, min2 := math.MaxInt, math.MaxInt
    for _, v := range nums {
        if v <= min1 {
            min1 = v
        } else if v <= min2 {
            min2 = v
        } else {
            return true
        }
    }
    return false
}
