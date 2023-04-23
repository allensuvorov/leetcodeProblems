func characterReplacement(s string, k int) int {
    count := map[byte]int{} // sliding collection
    res := 0
    maxF := 0 // all string max friquency
    l := 0
    for r := range s {
        count[s[r]]++
        if maxF < count[s[r]] {
            maxF = count[s[r]]
        }

        for r - l + 1 > k + maxF { // window
            count[s[l]]--
            l++
        }

        if res < r - l + 1 {
            res = r - l + 1
        }
    }
    return res
}
