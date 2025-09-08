func getNoZeroIntegers(n int) []int {
    for a := 1; a < n; a++ {
        if isNoZero(a) && isNoZero(n - a) {
            return []int{a, n-a}
        }
    }
    return nil
}

func isNoZero(a int) bool {
    for x := a; x != 0; x = x / 10 {
        d := x % 10
        if d % 10 == 0 {
            return false
        }
    }
    return true
}
