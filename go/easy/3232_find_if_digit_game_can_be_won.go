func canAliceWin(nums []int) bool {
    var sum int
    for _, v := range nums {
        if v > 9 {
            sum += v
        } else {
            sum -= v
        }
    }
    return sum != 0
}
