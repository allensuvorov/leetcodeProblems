func longestCommonPrefix(strs []string) string {
    ans := []byte(strs[0])
    for i := 1; i < len(strs); i++ {
        j := 0
        for j < len(ans) && j < len(strs[i]) && ans[j] == strs[i][j] {
            j++
        }
        ans = ans[:j]
    }
    return string(ans)
}
