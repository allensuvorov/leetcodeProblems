func toLowerCase(s string) string {
    shift := 'a' - 'A'
    res := make([]byte, len(s))
    for i, v := range s {
        if v >= 'A' && v <= 'Z' {
            res[i] = byte(v+shift)
        } else {
            res[i] = byte(v)
        }
    }
    return string(res)
}
