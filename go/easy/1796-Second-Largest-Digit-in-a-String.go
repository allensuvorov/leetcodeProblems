func secondHighest(s string) int {
    first, second := -1, -1
    for _, c := range s {
        if unicode.IsDigit(c) {
            num := int(c-'0')
            if num > first {
                second = first
                first = num
            } else if num > second && num < first {
                second = num
            }
        }
    }
    return second
}
