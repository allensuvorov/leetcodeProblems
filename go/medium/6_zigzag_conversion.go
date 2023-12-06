func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    
    res := make([]byte, 0, len(s))
    step := (numRows-1)*2
    midStep := step
    
    for r := 0; r < numRows; r++ {
        midStep = step - r*2
        for i := r; i < len(s); i += step {
            res = append(res, s[i])
            if r != 0 && r != numRows - 1 && i + midStep < len(s) {
                res = append(res, s[i+midStep])
            }
        }
    }    
    return string(res)
}
