func removeDigit(number string, digit byte) string {
    max := ""
    for i := range number {
        if number[i] == digit {
            l := number[:i]
            r := number[i + 1:]
            if max < l + r {
                max = l + r
            }
        }
    }
    return max
}
