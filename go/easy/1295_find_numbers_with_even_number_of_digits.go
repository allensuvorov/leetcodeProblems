func findNumbers(nums []int) int {
    count := 0
    for _, v := range nums {
        digitCount := 0
        for v > 0 {
            digitCount++
            v = v / 10
        }
        if digitCount % 2 == 0 {
            count++
        }
    }
    return count
}
