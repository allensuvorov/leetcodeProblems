func selfDividingNumbers(left int, right int) []int {
    ans := []int{}
    for n := left; n <= right; n++ {
        if isSelfDifiding(n) {
            ans = append(ans, n)
        }
    }
    return ans
}

func isSelfDifiding(n int) bool {
    m := n
    for m > 0 {
        d := m % 10
        if d == 0 {
            return false
        } else {
            if n % d != 0 {
                return false
            }
        }
        m /= 10
    }
    return true
}
