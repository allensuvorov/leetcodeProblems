func backspaceCompare(s string, t string) bool {
    i := len(s) - 1
    j := len(t) - 1
    sBackspaceCount := 0
    tBackspaceCount := 0
    for i >= 0 && j >= 0 {
        if s[i] == '#' {
            sBackspaceCount++
            i--
            continue
        } else {
            if sBackspaceCount > 0 {
                i--
                sBackspaceCount--
                continue
            }
        }
        if t[j] == '#' {
            tBackspaceCount++
            j--
            continue
        } else {
            if tBackspaceCount > 0 {
                j--
                tBackspaceCount--
                continue
            }
        }
        if i >= 0 && j >= 0 {
            if s[i] != t[j] {
                return false
            }
        }
        i--
        j--
    }

    if i < 0 && j < 0 {
        return true
    }
    fmt.Println("i < 0 || j < 0", i, j)
    return false
}
