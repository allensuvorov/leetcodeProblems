func clearDigits(s string) string {
    var b []rune
    for _, v := range s {
        if unicode.IsDigit(v) {
            b = b[:len(b)-1]
        } else {
            b = append(b, v)
        }
    }
    return string(b)
}

/*
[abc321]
[ab21]
[a1]
[]
*/
