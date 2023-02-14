func checkString(s string) bool {
    // look for a after first b
    i := 0
    j := 0
    foundB := false
    for i < len(s) {
        if s[i] == 'b' {
            j = i
            foundB = true
            break
        }
        i++
    }

    if foundB {
        for j < len(s) {
            if s[j] == 'a' {
                return false
            }
            j++
        }
    }

    return true
}
