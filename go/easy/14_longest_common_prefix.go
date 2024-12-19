func longestCommonPrefix(strs []string) string {
    lcp := strs[0]
    for _, v := range strs{
        i := 0
        for i < len(lcp) && i < len(v) && lcp[i] == v[i] {
            i++
        }
        lcp = lcp[:i]
    }
    return string(lcp)
}
