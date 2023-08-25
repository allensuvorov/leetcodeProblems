func lengthOfLongestSubstring(s string) int {
    res := 0
    l, r := 0, 0
    m := map[byte]int{}
    for r < len(s) {
        if v, ok := m[s[r]]; ok {
            // clean tail
            for i := l; i <= v; i++{
                delete(m, s[i])
            }
            l = v + 1
        } else {
            m[s[r]] = r
            if len(m) > res {
                res = len(m)
            }
            r++
        }
    }
    return res
}
