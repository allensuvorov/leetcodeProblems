func isPalindrome(x int) bool {
    if x == 0 {
        return true
    }
    
    if x < 0 || x % 10 == 0 {
        return false
    }
    
    return x == revDig(x)
}

func revDig(x int) int {
    y := 0
    for x > 0 {
        y *= 10
        y += x % 10
        x /= 10
    }
    return y
}
