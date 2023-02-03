func toLowerCase(s string) string {
    shift := 'a' - 'A'
    res := ""
    for _, v := range s {
        if v >= 'A' && v <= 'Z' {
            res = res + string(v+shift)
        } else {
            res = res + string(v)
        }
    }
    return res
}
