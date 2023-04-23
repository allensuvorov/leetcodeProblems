func characterReplacement(s string, k int) int {
    count := [26]int{} // sliding collection
    res := 0
    maxF := 0 // all string max friquency
    l := 0
    for r := range s {
        count[s[r] -'A']++
        if maxF < count[s[r] - 'A'] {
            maxF = count[s[r] - 'A']
        }

        for r - l + 1 > k + maxF { // window
            count[s[l] - 'A']--
            l++
        }

        if res < r - l + 1 {
            res = r - l + 1
        }
    }
    return res
}
