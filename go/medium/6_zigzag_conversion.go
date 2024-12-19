func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    res := make([]byte, 0, len(s))
    for r := 0; r < numRows; r++ {
        inc := (numRows-1) * 2
        for i := r; i < len(s); i += inc {
            res = append(res, s[i])
            if r > 0 && r < numRows - 1 && i + inc - r * 2 < len(s) {
                res = append(res, s[i + inc - r * 2])
            }
        }
    }    
    return string(res)
}

// time O(n)
// space O(n)

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    
    rows := make([]string, numRows)
    row := 0
    shift := -1
    for i := range s {
        rows[row] += string(s[i])
        if row == 0 || row == len(rows) - 1 {
            shift = -shift
        }
        row += shift
    }

    res := ""
    for _, row := range rows {
        res += row
    }
    return res
}
