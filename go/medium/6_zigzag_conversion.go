func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    ss := make([][]byte, numRows)
    r := 0
    step := 0
    for i := range s {
        ss[r] = append(ss[r], s[i]) 
        if r == 0 {
            step = 1
        }
        if r == numRows - 1 {
            step = -1
        }
        r = r + step
    }
    res := make([]byte, 0, len(s))
    for r := range ss {
        res = append(res, ss[r]...)
    }
    return string(res)
}
