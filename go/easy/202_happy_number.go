func isHappy(n int) bool {
    nums := map[int]bool{}
    for !nums[n] {
        nums[n] = true
        sum := 0
        for x := n; x > 0; x /= 10 {
            d := x % 10
            sum += d*d
        }
        n = sum
    }
    return n == 1
}
