func reverseStr(s string, k int) string {
    b := make([]byte, len(s))
    cur := 0
    i := 0
    
    for i < len(b){
        cur = i
        for j := cur + k - 1; j >= cur && i < len(b); j -- {
            if j < len(b) {
                b[i] = s[j]
                i ++
            }
        }
        cur = i
        for j := i; j < cur + k && i < len(b); j ++ {
            if j < len(b) {
                b[i] = s[j]
                i ++
            }
        }
    }
    return string(b)
}
