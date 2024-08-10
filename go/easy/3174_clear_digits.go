func clearDigits(s string) string {
    stack := []rune{}
    for _, v := range s {
        if unicode.IsLetter(v) {
            stack = append(stack, v)
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return string(stack)
}
