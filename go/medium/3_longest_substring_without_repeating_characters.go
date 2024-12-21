func lengthOfLongestSubstring(s string) int {
    set := make(map[byte]int)
    res := 0
    for l, r := 0, 0; r < len(s); r++ {
        set[s[r]]++

        for set[s[r]] > 1 {
            set[s[l]]--
            l++
        }
        res = max(res, r - l + 1)
    }
    return res
}
