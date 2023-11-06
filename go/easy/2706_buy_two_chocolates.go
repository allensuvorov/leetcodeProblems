func buyChoco(prices []int, money int) int {
    // find min1 and min2
    min1, min2 := 100, 100
    for _, v := range prices {
        if v < min2 {
            if v < min1 {
                min2 = min1
                min1 = v
            } else {
                min2 = v
            }
        }
    }
    leftOver := money - (min1 + min2)
    if leftOver < 0 {
        return money
    }
    return leftOver
}
