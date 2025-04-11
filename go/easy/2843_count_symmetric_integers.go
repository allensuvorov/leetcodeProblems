func countSymmetricIntegers(low int, high int) int {
    // run the range 1 - 10_000
    // read digits
    // check the condition
    res := 0
    for num := low; num <= high; num++ {
        digits := []int{}
        for x := num; x > 0; x /=10 {
            digits = append(digits, x % 10)
        }

        if len(digits) % 2 == 0 {
            l, r := 0, len(digits) - 1
            leftSum, rightSum := 0, 0
            for l < r {
                leftSum += digits[l]
                rightSum += digits[r]
                l++
                r--
            }
            if leftSum == rightSum {
                res++
            }
        }
    }
    return res
}
