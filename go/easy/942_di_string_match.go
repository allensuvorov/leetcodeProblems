func diStringMatch(s string) []int {
    ans := make([]int, 0, len(s) + 1)
    l, r := 0, cap(ans) - 1
    for _, v := range s {
        if v == 'I' {
            ans = append(ans, l)
            l++
        } else {
            ans = append(ans, r)
            r--
        }
    }
    return append(ans, r)
}
