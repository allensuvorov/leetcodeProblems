func findNumbers(nums []int) int {
    count := 0
    for _, v := range nums {
        if v >= 10 && v <= 99 ||
           v >= 1_000 && v <= 9_999 ||
           v >= 100_000 && v <= 999_999 {
            count++
           }
    }
    return count
}
