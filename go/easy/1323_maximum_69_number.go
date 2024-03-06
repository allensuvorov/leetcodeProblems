func maximum69Number (num int) int {
    rev := 0
    res := 0
    for tmp := num; tmp > 0; tmp /= 10 {
        rev = rev * 10 + tmp % 10
    }

    firstSixFound := false
    for tmp := rev; tmp > 0; tmp /= 10 {
        d := tmp % 10
        if d == 6 && !firstSixFound {
            d = 9
            firstSixFound = true
        }
        res = res * 10 + d
    }
    return res
}
