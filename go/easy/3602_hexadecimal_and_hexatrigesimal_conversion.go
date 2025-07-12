func concatHex36(n int) string {
    hd := strings.ToUpper((strconv.FormatInt(int64(n*n), 16)))
    ht := strings.ToUpper((strconv.FormatInt(int64(n*n*n), 36)))
    return hd + ht
}


// custom formatInt() function:
func concatHex36(n int) string {
    hd := formatInt(n*n, 16)
    ht := formatInt(n*n*n, 36)
    return hd + ht
}

func formatInt(n, base int) string {
    s := ""
    for a := n; a > 0; a = a / base {
        d := a % base
        c := string(byte(d) + '0')
        if d > 9 {
            c = string(byte(d-10) + 'A')
        }
        s = c + s
    }
    return s
}
