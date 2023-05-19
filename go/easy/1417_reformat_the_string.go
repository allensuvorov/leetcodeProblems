func reformat(s string) string {
    sl := []rune(s)
    if len(s) == 1 {
        return s
    }
    if !isPossible(s) {
        return ""
    }
    p := 0 // pointer to opposite search start
    for i := 1; i < len(s); i++ {
        if unicode.IsNumber(sl[i]) == unicode.IsNumber(sl[i-1]) {
            if p == 0 {
                p = i + 1
            }
            for j := p; j < len(s); j ++ {
                if unicode.IsNumber(sl[j]) != unicode.IsNumber(sl[i-1]) {
                    sl[i], sl[j] = sl[j], sl[i]
                    p = j
                    break
                }
            }
        }
    }
    return string(sl)
}

func isPossible(s string) bool {
    n := 0
    for _, v := range s {
        if unicode.IsNumber(v) {
            n++
        }
    }
    if len(s) % 2 == 0 {
        if len(s) / 2 != n {
            return false
        }
    } else {
        if len(s) / 2 != n && len(s) / 2 + 1 != n {
            return false
        }
    }
    return true
}
