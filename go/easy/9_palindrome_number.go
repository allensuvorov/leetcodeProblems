func isPalindrome(x int) bool {
    b := x
    y := 0
    for b > 0 {
        y = y * 10 + b % 10
        b = b / 10
    }
    return x == y
}
