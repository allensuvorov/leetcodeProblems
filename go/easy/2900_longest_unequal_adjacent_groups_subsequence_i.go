func getLongestSubsequence(words []string, groups []int) []string {
    var res []string
    lastGroup := -1
    for i, v := range groups {
        if v != lastGroup {
            res = append(res, words[i])
            lastGroup = v
        }
    }
    return res
}
