func longestPalindrome(s string) string {
    n := len(s)
    ans := ""
    for i := range s {
        for offset := range 2 {
            l := i
            r := i + offset
            for l >= 0 && r < n && s[l] == s[r] {
                cur := s[l:r+1]
                if len(cur) > len(ans) {
                    ans = cur
                }
                l--
                r++
            }
        }
    }
    return ans
}
