func removeDigit(number string, digit byte) string {
    target := 0
    for i := range number {
        if number[i] == digit {
            if i < len(number) - 1 && number[i] < number[i+1] {
                target = i
                break
            }
            target = i
        }
    }
    l := number[:target]
    r := number[target + 1:]
    return l + r
}
