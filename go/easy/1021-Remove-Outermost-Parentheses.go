func removeOuterParentheses(s string) string {
    res := make([]byte, 0, len(s))
    count := 0

    for i := range s {
        switch s[i]{
            case '(':
                if count != 0 {
                    res = append(res, s[i])
                }
                count++
            case ')':
                if count != 1 {
                    res = append(res, s[i])
                }
                count--
        }
    }
    return string(res)
}
