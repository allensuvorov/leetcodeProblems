func longestPalindrome(s string) int {
    hm := map[byte]int{}
    res := 0
    for i := range s {
        hm[s[i]]++
    }
    center := false
    for _, v := range hm {
        if v%2 == 0 {
            res +=v
        } else {
            center = true
            res +=v-1
        }
    }
    if center {
        return res + 1
    } else {
        return res
    }

}
