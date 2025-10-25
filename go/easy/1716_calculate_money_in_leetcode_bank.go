func totalMoney(n int) int {
    accountBalanc := 0
    depositAmount := 1
    weeklyIncrement := 0
    for range n {
        accountBalanc += weeklyIncrement + depositAmount
        if depositAmount % 7 == 0 {
            weeklyIncrement++
            depositAmount = 0
        }
        depositAmount++
    }
    return accountBalanc
}
