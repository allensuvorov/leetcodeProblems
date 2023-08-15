func plusOne(digits []int) []int {
    i := len(digits) - 1
    for i >= 0 && digits[i] == 9 {
        digits[i] = 0
        i--
    }
    if i < 0 {
        digits = append([]int{1}, digits...)
    } else {
        digits[i] += 1
    }
    return digits
}
