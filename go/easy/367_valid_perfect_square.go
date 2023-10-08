func isPerfectSquare(num int) bool {
    l := 1
    r := num
    for l <= r {
        m := l + (r-l)/2
        switch{
        case m*m == num:
            return true
        case m*m < num:
            l = m + 1
        case m*m > num:
            r = m - 1
        }
    }
    return false
}
