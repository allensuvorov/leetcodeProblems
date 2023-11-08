func arraySign(nums []int) int {
    sign := 1
    for _, v := range nums {
        if v == 0 {
            return 0
        }
        if v < 0 {
            sign = -sign
        }
    }
    return sign
}
