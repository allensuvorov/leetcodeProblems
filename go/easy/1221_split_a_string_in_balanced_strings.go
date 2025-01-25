func balancedStringSplit(s string) int {
    b := 0
    res := 0
    for _, v := range s {
        if v == 'R' {
            b++
        } else {
            b--
        }
        if b == 0 {
            res++
        }
    }
    return res
}
