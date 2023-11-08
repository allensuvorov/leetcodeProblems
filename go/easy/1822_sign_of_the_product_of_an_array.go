func arraySign(nums []int) int {
    sign := false // positive
    for _, v := range nums {
        if v < 0 {
            sign = sign != true
        }
        if v == 0 {
            return 0
        }
    }
    if sign {
        return -1
    }
    return 1
}
