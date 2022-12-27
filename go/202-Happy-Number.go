func isHappy(n int) bool {
    slow, fast := n, n; 
    for {
        slow = digitSquareSum(slow)
        fast = digitSquareSum(fast)
        fast = digitSquareSum(fast)
        if slow == fast {
            break
        }
    }
    if slow == 1 {
        return true
    } else {
        return false
    }
}

func digitSquareSum(n int) int{
    var sum, d int 
    for {
        d = n%10
        sum += d*d
        if n < 10 {
            return sum
        }
        n /= 10
    }
}
