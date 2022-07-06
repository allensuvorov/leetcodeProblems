func plusOne(digits []int) []int {
    var carry int = 1
    for i:=len(digits)-1; i>=0; i-- {
        if digits[i] < 9 {
            digits[i] +=carry
            return digits
        }
        digits[i] = 0
    }
    digits = append ([]int{1}, digits...)
    return digits
}
