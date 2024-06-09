func backspaceCompare(s string, t string) bool {
    i := len(s) - 1
    j := len(t) - 1
    skipS := 0
    skipT := 0
    for i >= 0 || j >= 0 {
        for i >= 0 {
            if s[i] == '#' {
                skipS++
            } else if skipS > 0 {
                skipS--
            } else {
                break
            }
            i--
        }
        for j >= 0 {
            if t[j] == '#' {
                skipT++
            } else if skipT > 0 {
                skipT--
            } else {
                break
            }
            j--
        }

        if i >= 0 && j >= 0 && s[i] != t[j] {
            return false
        }
        
        if (i>=0) != (j>=0) {
            return false
        }
        i--
        j--

    }
    return true
}
