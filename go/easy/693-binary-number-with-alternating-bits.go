func hasAlternatingBits(n int) bool {
    var cur int = 0
    var prev int = -1 
    for n != 0 {
        cur = n%2
        if cur == prev {return false}
        n = n/2
        prev = cur
    }
    return true
}
