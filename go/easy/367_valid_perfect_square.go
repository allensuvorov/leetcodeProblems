func isPerfectSquare(num int) bool {
    l := 0
    r := num + 1

    for l < r {
        m := l + (r - l) / 2
        if m*m >= num {
            r = m
        } else {
            l = m + 1
        }
    }
    return l*l == num
}
