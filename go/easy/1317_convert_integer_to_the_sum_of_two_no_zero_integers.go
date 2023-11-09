func getNoZeroIntegers(n int) []int {
    a := 1
    m := n
    level := 1
    carry := 0
    if m > 9 {
        for a = 0; m / 10 != 0; m /= 10 {
            d := m % 10 - carry
            if d < 2 {
                a += 2 * level
                carry = 1
            } else {
                a += 1 * level
                carry = 0
            }
            level *= 10
        }
    }
    return []int{a, n - a}
}
