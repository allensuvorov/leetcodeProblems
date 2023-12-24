func removeStars(s string) string {
    stack := make([]byte, 0, len(s))
    for i := range s {
        if s[i] == '*' {
            stack = stack[:len(stack) - 1] // pop
        } else {
            stack = append(stack, s[i]) // push
        }
    }
    return string(stack)
}

func removeStars(s string) string {
    // reverse solution
    rev := make([]byte, 0, len(s))
    count := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '*' {
            count++
        } else {
            if count > 0 {
                count --
            } else {
                rev = append(rev, s[i])
            }
        }
    }
    res := make([]byte, 0, len(rev))
    for i := len(rev) - 1; i >= 0; i-- {
        res = append(res, rev[i])
    }
    return string(res)
}
