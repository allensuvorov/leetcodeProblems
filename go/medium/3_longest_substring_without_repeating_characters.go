func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]bool)
    res := 0
    for l, r := 0, 0; r < len(s); r++ {
        for set[s[r]] {
            set[s[l]] = false
            l++
        }
        set[s[r]] = true
        res = max(res, r - l + 1)
    }
    return res
}
