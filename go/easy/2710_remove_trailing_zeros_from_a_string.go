func removeTrailingZeros(num string) string {
    count := 0
    for i := len(num) - 1; i >= 0 && num[i] == '0' ; i-- {
        count++
    }
    return num[:len(num)-count]
}
