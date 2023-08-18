package easy

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    cf := [26]int{}
    for i := 0; i < len(s); i++ {
        cf[s[i]-'a']++
        cf[t[i]-'a']--
    }
    for _, v := range cf {
        if v != 0 {
            return false
        }
    }
    return true
}
