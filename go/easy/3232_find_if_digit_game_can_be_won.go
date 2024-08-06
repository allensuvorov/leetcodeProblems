func canAliceWin(nums []int) bool {
    var balance int
    for _, v := range nums {
        if v > 9 {
            balance += v
        } else {
            balance -= v
        }
    }
    return balance != 0
}
